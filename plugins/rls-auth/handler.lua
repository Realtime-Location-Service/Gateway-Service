local BasePlugin = require "kong.plugins.base_plugin"
local constants = require "kong.plugins.rls-auth.constants"
local httpStatus = require "kong.plugins.http_status"
local http = require "kong.plugins.http"
local cache = require "kong.plugins.cache"

local kong = kong
local AuthHandler = BasePlugin:extend()

AuthHandler.PRIORITY = constants.PRIORITY
AuthHandler.VERSION = constants.VERSION

function AuthHandler:new()
  AuthHandler.super.new(self, constants.PLUGIN_NAME)
end

function AuthHandler:access(conf)
  AuthHandler.super.access(self)
  local app_key = kong.request.get_header(constants.AUTH_HEADER)

  if not app_key or app_key == "" then
    return kong.response.exit(httpStatus.BADREQUEST, { message = "Missing AppKey!" })
  end

  local headers = {[constants.AUTH_HEADER]=app_key,[":method"]="GET"}
  local credential, err = cache.Get(app_key, {ttl = conf.request.cacheTTL},
                                  resolveDomain, conf, headers, http.Request)
  if err then
    cache.Forget(app_key)
    return kong.response.exit(httpStatus.SERVER_ERROR, { message = "Error happend while resolving domain!"})
  end

  if credential.status_code ~= httpStatus.OK or not credential.domain then
    return kong.response.exit(httpStatus.UNAUTHORIZED, { message = "You are not Authorized!"})
  end

  ngx.req.set_header(constants.RLS_REFERRER, credential.domain)

end

function resolveDomain(conf)
    if not conf or type(conf) ~= "table" or table.getn(conf) < 3 then
      return {status_code = "500", domain = nil}
    end

    local headers = conf[2]
    local http_req_func = conf[3]
    conf = conf[1]

    local status_code, body = http_req_func(conf.request.authURL, headers, conf.request.authTimeout)

    if not body or not body.domain then
      return {status_code = status_code, domain = nil}
    end

    return {status_code = status_code, domain = body.domain}
end


return AuthHandler

local BasePlugin = require "kong.plugins.base_plugin"
local constants = require "kong.plugins.rls-auth.constants"
local httpStatus = require "kong.plugins.http_status"
local http = require "kong.plugins.http"

local kong = kong
local AuthHandler = BasePlugin:extend()

AuthHandler.PRIORITY = constants.PRIORITY
AuthHandler.VERSION = constants.VERSION

function AuthHandler:new()
  AuthHandler.super.new(self, constants.PLUGIN_NAME)
end


function AuthHandler:access(theConf)
  AuthHandler.super.access(self)
  local app_key = kong.request.get_header(constants.AUTH_HEADER)

  if not app_key or app_key == "" then
    return kong.response.exit(httpStatus.BADREQUEST, { message = "Missing AppKey!" })
  end

  local headers = {[constants.AUTH_HEADER]=app_key,[":method"]="GET"}

  local status_code, body = http.Request(theConf.request.authURL, headers, theConf.request.authTimeout)
  if status_code ~= tostring(httpStatus.OK) then
    return kong.response.exit(httpStatus.UNAUTHORIZED, { message = "You are not Authorized!"})
  end

  if body ~= nil then
    ngx.req.set_header(constants.RLS_REFERRER, body.domain)
  end

end


return AuthHandler

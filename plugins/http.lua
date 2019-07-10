local http_request = require "http.request"
local http_headers = require "http.headers"
local httpStatus = require "kong.plugins.http_status"
local cjson = require "cjson"


function Request(url, headers, timeout)
  local body = "{}"

  if not url or url == "" then
    return httpStatus.BADREQUEST, cjson.decode(body)
  end

  local req = http_request.new_from_uri(url)
  for k, v in pairs(headers) do
    req.headers:upsert(k, v, false)
  end

  local headers, stream = req:go(timeout)
  body = stream:get_body_as_string(timeout)
  stream:shutdown()

  if not body or not string.find(headers:get("content-type"), "application/json") then
    body = "{}"
  end
  return headers:get ":status", cjson.decode(body)
end


return {
  Request = Request
}

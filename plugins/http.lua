local http = require("socket.http")
local httpStatus = require "kong.plugins.http_status"
local cjson = require "cjson"
local ltn12 = require "ltn12"

function Request(url, headers, timeout)
  local res = {}

  if not url or url == "" then
    return httpStatus.BADREQUEST, {}
  end

  local one, code = http.request{
    url = url,
    headers = headers,
    sink = ltn12.sink.table(res),
    create=function()
      local req_sock = socket.tcp()
      req_sock:settimeout(timeout, 't')
      return req_sock
    end
  }

  res = table.concat(res)

  return code, cjson.decode(res)
end


return {
  Request = Request
}

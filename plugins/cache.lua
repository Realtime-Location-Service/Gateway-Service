local kong =  kong

local CACHE_PREFIX = "API_KEYS."
local DEFAUTL_TTL = 300 -- TTL in second

local function Get(key, opts, cb_func, ...)
  opts = opts or {ttl = DEFAUTL_TTL}
  local cb_func_args = {...}

  return kong.cache:get(CACHE_PREFIX .. key, opts, cb_func, cb_func_args)
end


local function Forget(key)
  key = CACHE_PREFIX .. key
  
  local ttl, err, value = kong.cache:probe(key)
  if ttl then
    kong.cache:invalidate(key)
  end
end


return {
  Get = Get,
  Forget = Forget
}

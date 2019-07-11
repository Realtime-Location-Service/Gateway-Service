local typedefs = require "kong.db.schema.typedefs"
local constants = require "kong.plugins.rls-auth.constants"

local required_string = {
  type = "string",
  required = true,
}

local required_int = {
  type = "integer",
  required = true,
}

return {
  name = constants.PLUGIN_NAME,
  fields = {
    {
      -- this plugin will only be applied to Services or Routes
      consumer = typedefs.no_consumer
    },
    {
      config = {
        type = "record",
        fields = {
          {
            request = {
              type = "record",
              fields = {
                {
                  authURL = required_string,
                },
                {
                  authTimeout = required_int,
                }
              }
            }
          },
        }
      }
    }
  },
}

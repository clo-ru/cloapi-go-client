# Downstream spec patches. GOAL: shrink toward empty by fixing each defect at the
# API source. Nested tags, operationId cleanup, and servers are now correct upstream,
# so those rules were removed. The rules below patch defects still pending upstream.

# Force OpenAPI 3.0.0 for oapi-codegen compatibility (spec ships 3.1.0).
.openapi = "3.0.0" |

# Declare bearer auth.
# Pending upstream: the API does not yet emit a securityScheme.
# Tokens are UUIDs, not JWTs — no bearerFormat is claimed.
. + {"security": [{"bearerAuth": []}]} |
.components.securitySchemes = {"bearerAuth": {"type": "http", "scheme": "bearer"}} |

# Add a 400 response where an operation documents no 4xx.
(.paths[]? | .[]? | select(has("responses")) | .responses) |= (
    if (keys | map(startswith("4")) | any | not)
    then . + {"400": {"description": "Bad Request"}}
    else .
    end
) |

# Define the ApiError schema.
# Pending upstream: real error bodies differ from this shape, so these fields
# currently decode empty (IsNotFound still works — Code is seeded from HTTP status).
.components.schemas += {
  "ApiError": {
    "type": "object",
    "properties": {
      "message": { "type": "string" },
      "code": { "type": "integer" },
      "description": { "type": "string" }
    },
    "required": ["message", "code", "description"]
  }
} |

# Point every 4xx/5xx response at ApiError.
(.paths[]? | .[]? | .responses? | select(. != null)) |= with_entries(
  if (.key | tonumber? // 0) >= 400 then
    .value.content."application/json".schema = { "$ref": "#/components/schemas/ApiError" }
  else
    .
  end
)
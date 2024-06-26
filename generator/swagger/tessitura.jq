# add default response for error messages, referencing the ErrorMessage type
.paths |= map_values(
    . | map_values(
        .responses += {"default": 
                        {"description": "Error", 
                         "schema": {
                            "$ref": "#/definitions/ErrorMessage"
                            }
                        }
                    }
        )
    ) | 
# remove broken specs
    delpaths(
        [
            ["paths","/Batch"],
            ["paths","/Custom/{resourceName}"],
            ["paths","/Custom/{resourceName}/{id}"]
    ]) |
# fix duplicate enum
 .definitions.Response.properties.StatusCode.enum |= unique |
# make date-time nullable
.definitions |= map_values(
    .properties |= map_values( 
        select(.format == "date-time") += {"x-nullable":true}
    )
)
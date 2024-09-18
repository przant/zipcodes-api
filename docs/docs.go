// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/zipcodes/{zipcode}": {
            "get": {
                "description": "get info by zipcode",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Zipcode Information",
                "operationId": "get-info-by-zipcode",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Zipcode value",
                        "name": "zipcode",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Zipcode"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Zipcode": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "county": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "state_abbr": {
                    "type": "string"
                },
                "state_fips": {
                    "type": "string"
                },
                "zipcode": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:20790",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Zipcodes API",
	Description:      "Simple API for fetching US zipcodes and their related\ninformation like state, county, city,  and son on",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

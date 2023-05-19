// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/players": {
            "post": {
                "description": "プレイヤーを作成する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "players"
                ],
                "summary": "プレイヤーを作成する",
                "parameters": [
                    {
                        "description": "User information",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal.createPlayerRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/internal.createPlayerResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "internal.createPlayerRequest": {
            "type": "object",
            "required": [
                "at_bats",
                "hits",
                "home_runs",
                "name",
                "runs_batted_in",
                "uniform_number",
                "walks"
            ],
            "properties": {
                "at_bats": {
                    "type": "integer"
                },
                "hits": {
                    "type": "integer"
                },
                "home_runs": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "runs_batted_in": {
                    "type": "integer"
                },
                "uniform_number": {
                    "type": "integer"
                },
                "walks": {
                    "type": "integer"
                }
            }
        },
        "internal.createPlayerResponse": {
            "type": "object",
            "properties": {
                "at_bats": {
                    "type": "integer"
                },
                "hits": {
                    "type": "integer"
                },
                "home_runs": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "runs_batted_in": {
                    "type": "integer"
                },
                "uniform_number": {
                    "type": "integer"
                },
                "walks": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

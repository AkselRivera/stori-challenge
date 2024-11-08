// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "moralesaksel@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health": {
            "get": {
                "description": "Check if the service is up",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "The service is up",
                        "schema": {
                            "$ref": "#/definitions/health.health"
                        }
                    }
                }
            }
        },
        "/user/{user_id}/balance": {
            "get": {
                "description": "Retrieve the balance of a user with optional date range.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "Get user balance",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "2006-01-02T15:04:05Z07:00",
                        "description": "Start date in RFC3339 format",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "2006-01-02T15:04:05Z07:00",
                        "description": "End date in RFC3339 format",
                        "name": "to",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response with user balance",
                        "schema": {
                            "$ref": "#/definitions/domain.UserBalance"
                        }
                    },
                    "400": {
                        "description": "Error response for invalid input",
                        "schema": {
                            "$ref": "#/definitions/domain.CustomError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.CustomError": {
            "description": "An error that includes a specific code and a message with more details.",
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "domain.UserBalance": {
            "description": "User balance details including balance, total debits and total credits",
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number",
                    "example": 25
                },
                "total_credits": {
                    "type": "number",
                    "example": 15
                },
                "total_debits": {
                    "type": "number",
                    "example": 10
                }
            }
        },
        "health.health": {
            "description": "Detalles del estado de salud de la aplicación.",
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "ok"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Balance Service | API Docs",
	Description:      "This is a balance service for Stori Challenge",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

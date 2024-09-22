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
        "/api/v1/accounts": {
            "get": {
                "description": "get account state",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Account State",
                "operationId": "get-account-state",
                "parameters": [
                    {
                        "type": "string",
                        "description": "address",
                        "name": "addr",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/bets": {
            "get": {
                "description": "get active bets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Active Bets",
                "operationId": "get-active-bets",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/transactions": {
            "get": {
                "description": "get status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Transactions from the TON Blockchain",
                "operationId": "get-transactions",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Address",
                        "name": "addr",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Hash",
                        "name": "hash",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Lt",
                        "name": "lt",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

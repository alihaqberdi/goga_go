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
        "/api/client/tenders/{tender_id}/award/{id}": {
            "post": {
                "description": "Award a bid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bids"
                ],
                "summary": "Award a bid",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Tender ID",
                        "name": "tender_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Bid ID",
                        "name": "id",
                        "in": "path",
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
        "/api/client/tenders/{tender_id}/bids": {
            "get": {
                "description": "Get list of bids",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bids"
                ],
                "summary": "Get list of bids",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Tender ID",
                        "name": "tender_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.BidList"
                        }
                    }
                }
            }
        },
        "/api/contractor/bids/{tender_id}/bid/{id}": {
            "delete": {
                "description": "Delete a bid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bids"
                ],
                "summary": "Delete a bid",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Bid ID",
                        "name": "id",
                        "in": "path",
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
        "/api/contractor/tenders/{tender_id}/bid": {
            "post": {
                "description": "Create a new bid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bids"
                ],
                "summary": "Create a new bid",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Tender ID",
                        "name": "tender_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Bid object",
                        "name": "bid",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.BidCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.BidList"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.AuthRes"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "Register",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.Register"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.AuthRes"
                        }
                    }
                }
            }
        },
        "/users/{id}/bids": {
            "get": {
                "description": "Get user bids",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bids"
                ],
                "summary": "Get user bids",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.BidList"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.AuthRes": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/dtos.User"
                }
            }
        },
        "dtos.BidCreate": {
            "type": "object",
            "properties": {
                "comments": {
                    "type": "string",
                    "example": "This is a comment"
                },
                "contractor_id": {
                    "type": "integer",
                    "example": 1
                },
                "delivery_time": {
                    "type": "string",
                    "example": "12"
                },
                "price": {
                    "type": "number",
                    "example": 100
                },
                "status": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/types.BidStatus"
                        }
                    ],
                    "example": "pending"
                },
                "tender_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "dtos.BidList": {
            "type": "object",
            "properties": {
                "comments": {
                    "type": "string",
                    "example": "This is a comment"
                },
                "contractor_id": {
                    "type": "integer",
                    "example": 1
                },
                "delivery_time": {
                    "type": "string",
                    "example": "12"
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number",
                    "example": 100
                },
                "status": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/types.BidStatus"
                        }
                    ],
                    "example": "pending"
                },
                "tender_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "dtos.Login": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dtos.Register": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/types.UserRole"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dtos.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "role": {
                    "$ref": "#/definitions/types.UserRole"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "types.BidStatus": {
            "type": "string",
            "enum": [
                "pending",
                "awarded"
            ],
            "x-enum-varnames": [
                "BidStatusPending",
                "BidStatusAwarded"
            ]
        },
        "types.UserRole": {
            "type": "string",
            "enum": [
                "client",
                "contractor"
            ],
            "x-enum-varnames": [
                "UserRoleClient",
                "UserRoleContractor"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Go API Example",
	Description:      "API documentation for the Go application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

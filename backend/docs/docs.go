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
            "name": "Mark Raiter",
            "email": "raitermark@proton.me"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.LoginUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    }
                }
            }
        },
        "/logout": {
            "get": {
                "description": "Logout",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Logout",
                "operationId": "logout",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "create account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "SignUp",
                "operationId": "create-account",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateUserReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    }
                }
            }
        },
        "/ws/create-room": {
            "post": {
                "description": "create room",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Websocket"
                ],
                "summary": "CreateRoom",
                "operationId": "create-room",
                "parameters": [
                    {
                        "description": "room info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/websocket.CreateRoomReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    }
                }
            }
        },
        "/ws/get-clients": {
            "get": {
                "description": "get all clients in the room",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Websocket"
                ],
                "summary": "GetClients",
                "operationId": "get-clients",
                "parameters": [
                    {
                        "type": "string",
                        "description": "room_id",
                        "name": "room_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    }
                }
            }
        },
        "/ws/get-rooms": {
            "get": {
                "description": "get all rooms",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Websocket"
                ],
                "summary": "GetRooms",
                "operationId": "get-rooms",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    }
                }
            }
        },
        "/ws/join-room": {
            "get": {
                "description": "join room",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Websocket"
                ],
                "summary": "JoinRoom",
                "operationId": "join-room",
                "parameters": [
                    {
                        "type": "string",
                        "description": "room_id",
                        "name": "room_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "user.CreateUserReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.LoginUserReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "util.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "response message"
                }
            }
        },
        "websocket.CreateRoomReq": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "CHAT APP",
	Description:      "Docs for chat app backend API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
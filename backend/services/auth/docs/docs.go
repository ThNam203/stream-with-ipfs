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
            "name": "Nam Huynh",
            "email": "hthnam203@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Authenticate user with email and password",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Log in a user",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "userCredentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.logInForm"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "headers": {
                            "accessToken": {
                                "type": "string",
                                "description": "Access Token"
                            },
                            "refreshToken": {
                                "type": "string",
                                "description": "Refresh Token"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Username or password is not correct",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "description": "Register a new user with username, email, and password\nOn success, redirect user to index page and set refresh and access token in cookie",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Sign up a new user",
                "parameters": [
                    {
                        "description": "User registration data",
                        "name": "userForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.signUpForm"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "headers": {
                            "accessToken": {
                                "type": "string",
                                "description": "Access Token"
                            },
                            "refreshToken": {
                                "type": "string",
                                "description": "Refresh Token"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/verify": {
            "get": {
                "description": "Verifies a user's email address with the provided token",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Verify user email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email verification token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Return a Email verification complete! string",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Verify token expired or invalid.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "An error occurred while verifying the user.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.HTTPErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 500
                },
                "message": {
                    "type": "string",
                    "example": "internal server error"
                }
            }
        },
        "api.logInForm": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "hthnam203@gmail.com"
                },
                "password": {
                    "type": "string",
                    "maxLength": 72,
                    "minLength": 8,
                    "example": "123123123"
                }
            }
        },
        "api.signUpForm": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "hthnam203@gmail.com"
                },
                "password": {
                    "type": "string",
                    "maxLength": 72,
                    "minLength": 8,
                    "example": "123123123"
                },
                "username": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 6,
                    "example": "sen1or"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "localhost:8000",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Let's Live API",
	Description:      "The server API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
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
            "name": "Капитанов Даниил",
            "url": "https://vk.com/poophead27",
            "email": "kdanil01@mail.ru"
        },
        "license": {
            "name": "None",
            "url": "None"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v2/user/boards/": {
            "get": {
                "description": "Выводит и созданные им доски и те, в которых он гость. Работает только для авторизированного пользователя.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "boards"
                ],
                "summary": "Вывести все доски текущего пользователя",
                "responses": {
                    "200": {
                        "description": "Пользователь и его доски",
                        "schema": {
                            "$ref": "#/definitions/doc_structs.UserBoardsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperrors.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/apperrors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apperrors.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/login/": {
            "post": {
                "description": "Для этого использует сессии",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Войти в систему",
                "parameters": [
                    {
                        "description": "Эл. почта и логин пользователя",
                        "name": "authData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AuthInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Объект пользователя",
                        "schema": {
                            "$ref": "#/definitions/doc_structs.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperrors.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/apperrors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apperrors.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/logout/": {
            "delete": {
                "description": "Удаляет текущую сессию пользователя. Может сделать только авторизированный пользователь.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Выйти из системы",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Токен сессии",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "no content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperrors.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/apperrors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apperrors.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/signup/": {
            "post": {
                "description": "Также вводит пользователя в систему",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Зарегистрировать нового пользователя",
                "parameters": [
                    {
                        "description": "Эл. почта и логин пользователя",
                        "name": "authData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AuthInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Объект пользователя",
                        "schema": {
                            "$ref": "#/definitions/doc_structs.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperrors.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/apperrors.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/apperrors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apperrors.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/verify": {
            "get": {
                "description": "Узнать существует ли сессия текущего пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Подтвердить вход",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Токен сессии",
                        "name": "Cookie",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Объект пользователя",
                        "schema": {
                            "$ref": "#/definitions/doc_structs.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperrors.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/apperrors.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apperrors.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apperrors.ErrorResponse": {
            "type": "object",
            "properties": {
                "error_response": {
                    "type": "string"
                }
            }
        },
        "doc_structs.UserBoardsResponse": {
            "type": "object",
            "properties": {
                "boards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Board"
                    }
                },
                "user": {
                    "$ref": "#/definitions/entities.User"
                }
            }
        },
        "doc_structs.UserResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/entities.User"
                }
            }
        },
        "dto.AuthInfo": {
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
        "dto.UserInfo": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "entities.Board": {
            "type": "object",
            "properties": {
                "board_id": {
                    "type": "integer"
                },
                "guests": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.UserInfo"
                    }
                },
                "name": {
                    "type": "string"
                },
                "owner": {
                    "$ref": "#/definitions/dto.UserInfo"
                },
                "thumbnail_url": {
                    "type": "string"
                }
            }
        },
        "entities.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                },
                "thumbnail_url": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v2",
	Schemes:          []string{},
	Title:            "LA TABULA API",
	Description:      "Лучшее и единственно приложение, имитирующее Trello.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

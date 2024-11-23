// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/song": {
            "post": {
                "description": "Create a new song",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Songs"
                ],
                "summary": "Add a new song",
                "parameters": [
                    {
                        "description": "Song data",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.SongRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.SongInfoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse400"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse404"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse422"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse500"
                        }
                    }
                }
            }
        },
        "/song-verse": {
            "get": {
                "description": "Get song with pagination on verses",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Songs"
                ],
                "summary": "Get Songs With Verses",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Song filter",
                        "name": "song",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Group filter",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Verses filter",
                        "name": "verses",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.SongResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse400"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse404"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse500"
                        }
                    }
                }
            }
        },
        "/song/{id}": {
            "delete": {
                "description": "Delete songs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Songs"
                ],
                "summary": "Delete songs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "song id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.DeleteResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse400"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse500"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update songs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Songs"
                ],
                "summary": "Update songs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "song id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Song data",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UpdateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse400"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse500"
                        }
                    }
                }
            }
        },
        "/songs": {
            "get": {
                "description": "Get songs with filtr and pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Songs"
                ],
                "summary": "Get songs",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of items per page",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Group filter",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Song filter",
                        "name": "song",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "releaseDate filter",
                        "name": "releaseDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Text filter",
                        "name": "text",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Link filter",
                        "name": "link",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.SongsPaginationResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse404"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse500"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "requests.SongRequest": {
            "description": "Структура запроса для создания новой песни",
            "type": "object",
            "required": [
                "group",
                "song"
            ],
            "properties": {
                "group": {
                    "type": "string",
                    "minLength": 0,
                    "example": "Muse"
                },
                "song": {
                    "type": "string",
                    "minLength": 0,
                    "example": "Supermassive Black Hol"
                }
            }
        },
        "requests.UpdateRequest": {
            "type": "object",
            "required": [
                "group",
                "link",
                "releaseDate",
                "song",
                "text"
            ],
            "properties": {
                "group": {
                    "type": "string",
                    "minLength": 0,
                    "example": "Eminem"
                },
                "link": {
                    "type": "string",
                    "minLength": 0,
                    "example": "http://example.com"
                },
                "releaseDate": {
                    "type": "string",
                    "minLength": 0,
                    "example": "00.00.00"
                },
                "song": {
                    "type": "string",
                    "minLength": 0,
                    "example": "SOng"
                },
                "text": {
                    "type": "string",
                    "minLength": 0,
                    "example": "LaLala"
                }
            }
        },
        "responses.DeleteResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Delete succeeded"
                }
            }
        },
        "responses.ErrorResponse400": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "Bad Request - Invalid input data"
                }
            }
        },
        "responses.ErrorResponse404": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 404
                },
                "message": {
                    "type": "string",
                    "example": "Not Found - Song not found"
                }
            }
        },
        "responses.ErrorResponse422": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 422
                },
                "message": {
                    "type": "string",
                    "example": "Unprocessable Entity - Validation failed"
                }
            }
        },
        "responses.ErrorResponse500": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 500
                },
                "message": {
                    "type": "string",
                    "example": "Internal Server Error"
                }
            }
        },
        "responses.SongInfoResponse": {
            "description": "Response structure containing song information",
            "type": "object",
            "required": [
                "link",
                "releaseDate",
                "text"
            ],
            "properties": {
                "group": {
                    "type": "string",
                    "example": "Muse"
                },
                "link": {
                    "type": "string",
                    "minLength": 0,
                    "example": "https://www.youtube.com/watch?v=Xsp3_a-PMTw"
                },
                "releaseDate": {
                    "type": "string",
                    "minLength": 0,
                    "example": "16.07.2006"
                },
                "song": {
                    "type": "string",
                    "example": "Supermassive Black Hol"
                },
                "text": {
                    "type": "string",
                    "minLength": 0,
                    "example": "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight"
                }
            }
        },
        "responses.SongResponse": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
                "verses": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "responses.SongsPaginationResponse": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer",
                    "example": 10
                },
                "page": {
                    "type": "integer",
                    "example": 1
                },
                "songs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/responses.SongInfoResponse"
                    }
                },
                "total": {
                    "type": "integer",
                    "example": 31
                },
                "total_pages": {
                    "type": "integer",
                    "example": 4
                }
            }
        },
        "responses.UpdateResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Update succeeded"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "TZ",
	Description:      "Rest API Library",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

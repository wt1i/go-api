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
        "/api/v1/migrate": {
            "post": {
                "description": "migrate db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "migrate"
                ],
                "summary": "migrate db",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/news": {
            "post": {
                "description": "create news",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "News"
                ],
                "summary": "create an news",
                "parameters": [
                    {
                        "description": "news info",
                        "name": "request_data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/application.UpsertNewsReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.News"
                        }
                    }
                }
            }
        },
        "/api/v1/news/:news_id": {
            "get": {
                "description": "get news by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "News"
                ],
                "summary": "Show an news",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "News ID",
                        "name": "news_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.News"
                        }
                    }
                }
            },
            "post": {
                "description": "update news",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "News"
                ],
                "summary": "update an news",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "News ID",
                        "name": "news_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "news info",
                        "name": "request_data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/application.UpsertNewsReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "description": "remove news by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "News"
                ],
                "summary": "remove an news",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "News ID",
                        "name": "news_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.News"
                        }
                    }
                }
            }
        },
        "/api/v1/topic": {
            "get": {
                "description": "Show all topic",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "topic"
                ],
                "summary": "Show all topic",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Topic"
                        }
                    }
                }
            },
            "post": {
                "description": "create topic",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Topic"
                ],
                "summary": "create an topic",
                "parameters": [
                    {
                        "description": "topic info",
                        "name": "request_data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/application.UpsertTopicReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Topic"
                        }
                    }
                }
            }
        },
        "/api/v1/topic/:topic_id": {
            "get": {
                "description": "get topic by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "topic"
                ],
                "summary": "Show an topic",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "topic ID",
                        "name": "topic_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Topic"
                        }
                    }
                }
            },
            "post": {
                "description": "update topic",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Topic"
                ],
                "summary": "update an topic",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "topic ID",
                        "name": "topic_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "topic info",
                        "name": "request_data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/application.UpsertTopicReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "description": "remove topic by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Topic"
                ],
                "summary": "remove an topic",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "topic ID",
                        "name": "topic_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Topic"
                        }
                    }
                }
            }
        },
        "/functions": {
            "get": {
                "description": "show all news by status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "News"
                ],
                "summary": "Show all news",
                "parameters": [
                    {
                        "type": "string",
                        "description": "news's status exist draft|deleted|publish",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page, default is 1",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page size, default is 20",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "order by, default is id, enable multiple fields, example: ordering=-name,id",
                        "name": "ordering",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.News"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "application.UpsertNewsReq": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "slug": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "application.UpsertTopicReq": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                }
            }
        },
        "model.News": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "slug": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "topic": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Topic"
                    }
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.Topic": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "news": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.News"
                    }
                },
                "slug": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
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

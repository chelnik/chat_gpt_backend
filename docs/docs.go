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
            "name": "https://t.me/Sayzeks"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/query": {
            "post": {
                "description": "Работает без поддержания диалога",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gptResponse"
                ],
                "summary": "Вопрос-ответ",
                "parameters": [
                    {
                        "description": "Передача ключа доступа и запроса к ChatGPT",
                        "name": "q",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.SingleQuery"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.ChatResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.ChatCompletionChoice": {
            "type": "object",
            "properties": {
                "finish_reason": {
                    "type": "string"
                },
                "index": {
                    "type": "integer"
                },
                "message": {
                    "$ref": "#/definitions/domain.ChatCompletionMessage"
                }
            }
        },
        "domain.ChatCompletionMessage": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "domain.ChatResponse": {
            "type": "object",
            "properties": {
                "choices": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.ChatCompletionChoice"
                    }
                },
                "created": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "object": {
                    "type": "string"
                },
                "usage": {
                    "$ref": "#/definitions/domain.Usage"
                }
            }
        },
        "domain.SingleQuery": {
            "type": "object",
            "properties": {
                "inquiry": {
                    "type": "string"
                },
                "key": {
                    "type": "string"
                }
            }
        },
        "domain.Usage": {
            "type": "object",
            "properties": {
                "completion_tokens": {
                    "type": "integer"
                },
                "prompt_tokens": {
                    "type": "integer"
                },
                "total_tokens": {
                    "type": "integer"
                }
            }
        },
        "handlers.ErrResp": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "ChatGpt backend",
	Description:      "Backend приложения для общения с ChatGPT",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

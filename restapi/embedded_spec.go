// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "produces": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "From the todo list tutorial on goswagger.io",
    "title": "A Todo list application",
    "version": "1.0.0"
  },
  "paths": {
    "/consumers/{id}": {
      "get": {
        "tags": [
          "consumers"
        ],
        "operationId": "findConsumer",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/consumer"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/todos": {
      "get": {
        "tags": [
          "todos"
        ],
        "operationId": "findTodos",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list the todo operations",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/item"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "todos"
        ],
        "operationId": "addTodo",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/todos/{id}": {
      "get": {
        "tags": [
          "todos"
        ],
        "operationId": "getTodo",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "tags": [
          "todos"
        ],
        "operationId": "updateTodo",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "todos"
        ],
        "operationId": "deleteTodo",
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "consumer": {
      "type": "object",
      "required": [
        "key"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "key": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        },
        "statusCode": {
          "type": "integer",
          "format": "int64"
        },
        "statusText": {
          "type": "string"
        }
      }
    },
    "item": {
      "type": "object",
      "required": [
        "description"
      ],
      "properties": {
        "completed": {
          "type": "boolean"
        },
        "description": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        }
      }
    }
  },
  "securityDefinitions": {
    "basicAuth": {
      "description": "Consumer key and id pairs",
      "type": "basic"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "produces": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "From the todo list tutorial on goswagger.io",
    "title": "A Todo list application",
    "version": "1.0.0"
  },
  "paths": {
    "/consumers/{id}": {
      "get": {
        "tags": [
          "consumers"
        ],
        "operationId": "findConsumer",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/consumer"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/todos": {
      "get": {
        "tags": [
          "todos"
        ],
        "operationId": "findTodos",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list the todo operations",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/item"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "todos"
        ],
        "operationId": "addTodo",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/todos/{id}": {
      "get": {
        "tags": [
          "todos"
        ],
        "operationId": "getTodo",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "tags": [
          "todos"
        ],
        "operationId": "updateTodo",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "todos"
        ],
        "operationId": "deleteTodo",
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "consumer": {
      "type": "object",
      "required": [
        "key"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "key": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        },
        "statusCode": {
          "type": "integer",
          "format": "int64"
        },
        "statusText": {
          "type": "string"
        }
      }
    },
    "item": {
      "type": "object",
      "required": [
        "description"
      ],
      "properties": {
        "completed": {
          "type": "boolean"
        },
        "description": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        }
      }
    }
  },
  "securityDefinitions": {
    "basicAuth": {
      "description": "Consumer key and id pairs",
      "type": "basic"
    }
  }
}`))
}
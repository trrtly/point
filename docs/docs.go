// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "wenlong",
            "url": "http://yy-git.youyao99.com/youyao/point",
            "email": "wenlong.chen@youyaomedtech.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/point/activity": {
            "post": {
                "description": "添加事件积分",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "事件积分"
                ],
                "summary": "添加事件积分",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/activity.Create"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功返回值",
                        "schema": {
                            "$ref": "#/definitions/render.Response"
                        }
                    },
                    "400": {
                        "description": "失败返回值",
                        "schema": {
                            "$ref": "#/definitions/render.Response"
                        }
                    }
                }
            }
        },
        "/api/point/goods": {
            "post": {
                "description": "积分兑换商品",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "兑换商品"
                ],
                "summary": "积分兑换商品",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/goods.Create"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功返回值",
                        "schema": {
                            "$ref": "#/definitions/render.Response"
                        }
                    },
                    "400": {
                        "description": "失败返回值",
                        "schema": {
                            "$ref": "#/definitions/render.Response"
                        }
                    }
                }
            }
        },
        "/api/point/page": {
            "post": {
                "description": "添加页面浏览积分",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "页面浏览积分"
                ],
                "summary": "添加页面浏览积分",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/page.Create"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功返回值",
                        "schema": {
                            "$ref": "#/definitions/render.Response"
                        }
                    },
                    "400": {
                        "description": "失败返回值",
                        "schema": {
                            "$ref": "#/definitions/render.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "activity.Create": {
            "type": "object",
            "required": [
                "activityCode",
                "uid"
            ],
            "properties": {
                "activityCode": {
                    "description": "事件编号",
                    "type": "string"
                },
                "specialCode": {
                    "description": "特例编号",
                    "type": "string"
                },
                "specialVal": {
                    "description": "特例数值",
                    "type": "string"
                },
                "uid": {
                    "description": "用户 id",
                    "type": "integer"
                }
            }
        },
        "goods.Create": {
            "type": "object",
            "required": [
                "goodsNum",
                "goodsYyid",
                "uid"
            ],
            "properties": {
                "goodsNum": {
                    "description": "商品数量",
                    "type": "integer"
                },
                "goodsYyid": {
                    "description": "商品编号",
                    "type": "string"
                },
                "uid": {
                    "description": "用户 id",
                    "type": "integer"
                }
            }
        },
        "page.Create": {
            "type": "object",
            "required": [
                "uid",
                "uri"
            ],
            "properties": {
                "uid": {
                    "description": "用户id",
                    "type": "integer"
                },
                "uri": {
                    "description": "页面路径的 ` + "`" + `path` + "`" + ` 部分，例如：` + "`" + `http://api.youyao.com/user/point` + "`" + `，则 ` + "`" + `uri` + "`" + ` 为 ` + "`" + `user/point` + "`" + `",
                    "type": "string"
                }
            }
        },
        "render.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "` + "`" + `code` + "`" + ` 错误码\n全局错误码说明：\n` + "`" + `1001` + "`" + ` 用户不存在",
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "description": "` + "`" + `msg` + "`" + ` 错误信息",
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "优药积分系统 api",
	Description: "优药积分系统 api",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}

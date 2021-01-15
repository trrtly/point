# 项目概述

* 产品名称：优药积分系统 api
* 项目代号：point
* 官方地址：<http://yy-git.youyao99.com/youyao/point>

## 运行环境要求

* Golang 1.15.0+

## 开发环境部署/安装

请确保本地安装了 [go 开发环境](https://golang.org/doc/install)，并开启 [go module](https://learnku.com/articles/27401)模式。

### 基础安装

#### 1. 克隆源代码

克隆 `http://yy-git.youyao99.com/youyao/point.git` 源代码到本地：

```bash
git clone http://yy-git.youyao99.com/youyao/point.git
```

#### 2. 生成配置文件

```bash
cp .env.example .env
```

你可以根据情况修改 `.env` 文件里的内容，如数据库连接、缓存 等：

```env
# db
POINT_DATABASE_TYPE=mysql
POINT_DATABASE_USER=root
POINT_DATABASE_PASSWORD=root
POINT_DATABASE_HOST=127.0.0.1
POINT_DATABASE_NAME=youyao_qa
```

#### 3. 启动服务

```bash
go run cmd/server/main.go cmd/server/wire_gen.go cmd/server/inject_server.go cmd/server/inject_store.go cmd/server/inject_hashids.go
```

### 链接入口

* 首页地址：<http://localhost:8080/point/index.html>

至此, 安装完成 ^_^。

## 主要依赖列表

* web：[fiber](https://github.com/gofiber/fiber)
* 依赖注入：[wire](https://github.com/google/wire)
* 环境配置：[godotenv](https://github.com/joho/godotenv)
* swagger：[swag](https://github.com/swaggo/swag)
* 表单验证：[validator](https://github.com/go-playground/validator)
* orm：[gorm](https://gorm.io)

## 参考项目

* 布局规范：[Standard Go Project Layout](https://github.com/golang-standards/project-layout)
* 程序编码：[kubernetes](https://github.com/kubernetes/kubernetes)、[drone](https://github.com/drone/drone)

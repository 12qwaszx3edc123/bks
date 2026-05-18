# bks

> **Author**: 12qwaszx3edc123 (bks)  
> **License**: MIT License  
> **GitHub**: https://github.com/12qwaszx3edc123/bks

Go 微服务脚手架生成工具 — 一键生成完整的微服务项目结构，支持数据库字段自动映射

## 安装

```bash
go install github.com/12qwaszx3edc123/bks/cmd/mygen@v1.0.0
```

## 快速开始

```bash
# 生成名为 myshop 的项目，包含 user 和 order 模块
mygen --name myshop --modules user,order --db-name myshop_db

# 指定数据库连接，自动读取表字段生成 model/proto/server
mygen --name myproject --modules user,order,doctor \
  --db-host 127.0.0.1 --db-port 3306 \
  --db-user root --db-password root \
  --db-name mydb

# 使用外部模板
mygen --name myproject --modules user --template /path/to/templates
```

## 参数说明

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `--name` | `bks` | 项目名称（绝对路径或相对路径） |
| `--modules` | - | 模块列表，逗号分隔，如 `user,order,doctor`（必填） |
| `--bff` | `h5` | BFF 类型: `h5`, `web`, `applet`, `app` |
| `--db-host` | `127.0.0.1` | 数据库主机 |
| `--db-port` | `3306` | 数据库端口 |
| `--db-user` | `root` | 数据库用户名 |
| `--db-password` | `root` | 数据库密码 |
| `--db-name` | - | 数据库名称（填写后自动读取表字段） |
| `--template` | 内嵌模板 | 外部模板目录路径 |

## 生成的项目结构

```
<项目名>/
├── common/
│   ├── config/          # 配置管理（config.yaml / config.go / global.go）
│   ├── init/            # 初始化（MySQL / Redis / ES / gRPC）
│   └── model/           # 数据模型（每个模块一个文件）
├── proto/               # Protobuf 定义（每个模块一个目录）
├── <module>-server/     # 各模块的 gRPC 服务
│   ├── basic/cmd/       # 服务入口 main.go
│   └── server/          # gRPC 服务实现
├── bff<type>/           # BFF 层（h5/web/applet/app）
│   ├── basic/           # 基础配置（main / middlewares / router）
│   └── handler/         # 请求处理（api / request）
└── pkg/                 # 公共工具包（jwt / upload / cart / alipay / sendsms / ordersn）
```

## 核心特性

### 数据库字段自动映射

指定 `--db-name` 后，mygen 会自动连接数据库读取表结构，生成完整的字段映射：

- **Model** — GORM 模型，包含正确的字段类型和 JSON 标签
- **Proto** — Protobuf 消息定义，MySQL 类型自动转换为 Proto 类型
- **Server** — gRPC 服务方法
- **Handler** — BFF 请求处理
- **Request** — 请求/响应结构体

### 多模块支持

支持一次生成多个微服务模块，gRPC 端口从 50051 自动递增：

```bash
mygen --name myproject --modules user,order,doctor
# user-server   → :50051
# order-server  → :50052
# doctor-server → :50053
```

### 内嵌模板

模板通过 `//go:embed` 内嵌到二进制文件中，无需额外配置。也支持 `--template` 参数指定外部模板目录进行自定义。

### MySQL 类型自动转换

| MySQL 类型 | Go 类型 | Proto 类型 |
|-----------|---------|-----------|
| varchar / char / text | string | string |
| int / tinyint / smallint | int32 | int32 |
| bigint | int64 | int64 |
| float | float32 | float |
| double / decimal | float64 | double |
| date / datetime / timestamp | string | string |
| bit / bool | bool | bool |

## 模板文件

模板位于 `templates/bks/` 目录：

```
templates/bks/
├── bff/
│   ├── cmd/main.go.tmpl
│   ├── handler/api/handler.go.tmpl
│   ├── handler/api/upload.go.tmpl
│   ├── handler/request/request.go.tmpl
│   ├── handler/request/upload.go.tmpl
│   ├── middlewares/middlewares.go.tmpl
│   └── router/router.go.tmpl
├── common/
│   ├── config/config.go.tmpl
│   ├── config/config.yaml.tmpl
│   ├── config/global.go.tmpl
│   ├── init/init.go.tmpl
│   └── model/model.go.tmpl
├── pkg/
│   ├── alipay.go.tmpl
│   ├── cart.go.tmpl
│   ├── jwt.go.tmpl
│   ├── ordersn.go.tmpl
│   ├── sendsms.go.tmpl
│   └── upload.go.tmpl
├── proto/proto.tmpl
└── srv/
    ├── cmd/main.go.tmpl
    └── server/server.go.tmpl
```

## 后续步骤

生成项目后：

1. 修改 `common/config/config.yaml` 中的空值配置（MySQL / Redis / ES 等）
2. 运行 `protoc` 生成 Go 代码
3. 运行 `go mod tidy` 下载依赖
4. 启动服务

## 技术栈

生成的项目使用以下技术栈：

- **Web 框架**: Gin
- **ORM**: GORM + MySQL
- **缓存**: Redis (go-redis/v9)
- **搜索引擎**: ElasticSearch (olivere/elastic/v7)
- **RPC**: gRPC + Protobuf
- **配置**: Viper
- **认证**: JWT (dgrijalva/jwt-go)
- **支付**: Alipay (smartwalle/alipay)
- **短信**: 自定义短信服务
- **对象存储**: 七牛云 (qiniu/go-sdk)

## License

MIT License

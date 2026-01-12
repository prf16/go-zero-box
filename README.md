## 概念

### 概述

基于 go-zero 框架的 1.9.4 版本、goctl 1.9.2 版本.

开箱即用的 [go-zero](https://go-zero.dev) 示例，做了工程化的包装，包装后的框架包含 api、scheduler、queue、script 服务。

在熟悉 go-zero 框架过程中踩坑很多，衍生的 go-zero-box 框架也是相当于经验贴了，希望可以帮助开发者更快的度过熟悉期。

整理并且开源不易，后续还会持续编写 go-zero 框架和 go 语言新特性的使用心得，在使用过程中有任何问题，欢迎提 **issue**，如果您觉得项目对您有帮助，希望给个️ ⭐️ **star** ⭐️ ，十分感谢。

go-zero 是一个集成了各种工程实践的 web 和 rpc 框架，缩短从需求到上线的距离，通过弹性设计保障了大并发服务端的稳定性，经受了充分的实战检验。

### 特性

- go-zreo 框架 + goctl 代码生成工具 + wire 依赖注入 + Makefile 实现自动化编译；
- 做了一些符合Web应用框架的目录划分，使开发者代码阅读更便捷，减少开发成本；
- 支持多库，多redis；
- 提供了初始业务的 api 代码示例，包含登录、注册、鉴权等开箱即用的业务功能；
- 提供了 go-zero 支持的 Swagger api 文档服务，**访问/api/doc即可预览文档**；
- 增加了更多的 ORM API，支持复杂的数据操作业务；
- 引入 asynq 库，实现了队列、调度、脚本等功能；
- 引入 wire 库，实现了依赖注入功能；
- 通过 ```app [command]``` 的方式，在分布式场景下可以更方便地部署 api 、scheduler、queue 服务；
- go 语言新特性的优秀示例，比如通过范式实现的三元表达式 ```tools.Any(a == b, a, b)```；


    ···
    
    持续更新中，敬请期待。


### 安装开发工具

#### wire
```shell
# wire 是一个依赖注入工具，用于解决 Go 语言中依赖注入的问题。通过 wire，您可以轻松地管理应用程序的依赖关系，并确保它们在编译时进行注入。
$ go install github.com/google/wire/cmd/wire@latest

# 验证安装
$ wire
```

#### goctl 安装
```shell
# https://github.com/zeromicro/go-zero/releases/tag/tools%2Fgoctl%2Fv1.9.2
# goctl 是 go-zero 微服务框架下的代码生成工具。使用 goctl 可显著提升开发效率，让开发人员将时间重点放在业务开发上，其功能有：api服务生成、rpc服务生成、model代码生成、模板管理。
$ go install github.com/zeromicro/go-zero/tools/goctl@v1.9.2

# 验证安装
$ goctl --version
```

### Make 命令介绍

Makefile 文件描述了 Linux 系统下项目工程的编译规则，只需要一个 `make bild` 命令，整个工程就开始自动构建项目环境，不再需要手动执行大量的 `go build` 命令，Makefile 文件定义了一系列规则，指明了源文件的编译顺序、依赖关系、是否需要重新编译等，可以输入 `make help` 查看命令集。

```shell
# 生成api代码文件，并且更新 swagger 接口文档
$ make api

# 生成model代码文件
$ make model

# 生成wire代码文件
$ make wire

# 构建项目
$ make build env=dev|test|prod
```

### 代码结构

```text
.
├── app                             包含应用程序的主要代码
│   ├── api                         api 描述文件
│   ├── etc                         静态配置文件目录
│   ├── internal                    内部业务逻辑
│   ├── app.go                      应用程序的入口文件，定义了 api、scheduler、queue、script 服务。
│   ├── wire.go                     依赖注入文件
│   └── wire_gin.go                 依赖注入生成文件
├── build                           项目构建目录
│   ├── app                         应用构建后的目录  
│   │   ├── api                     接口文档目录     
│   │   ├── etc                     配置文件目录
│   │   └── app                     应用二进制文件
│   └── app.tar                     应用构建后的打包文件
├── deploy                          部署相关目录
│   ├── access                      示例图片                       
│   ├── goctl                       模板文件
│   └── sql                         初始化数据库SQL
├── pkg                             工具包
├── runtime                         项目运行时目录
├── vendor                          项目依赖包
├── .gitignore                      git 忽略文件
├── go.mod                          项目依赖管理文件
├── Makefile                        项目构建文件
└── README.md                       项目说明文件
```

### 环境要求
- Golang >= 1.20
- MySql >= 5.7


## 快速开始实例

### 初始化项目

#### 1. clone项目

![初始化项目](deploy/access/快速开始/1初始化项目.jpg)

#### 2. 打开项目

![打开项目](deploy/access/快速开始/2打开项目.jpg)

#### 3. 替换包名引用

![替换包名引用](deploy/access/快速开始/3替换包名引用.jpg)

#### 4. 启动go模块

![启动go模块](deploy/access/快速开始/4启动go模块.jpg)

#### 5. 安装gosdk

![下载依赖包](deploy/access/快速开始/5安装gosdk.png)

#### 6. 下载依赖包

![下载依赖包](deploy/access/快速开始/6下载依赖包.jpg)

#### 7. 启动redis

![启动redis](deploy/access/快速开始/7启动redis.jpg)

#### 8. 启动项目

![启动项目](deploy/access/快速开始/8启动项目.jpg)

#### 9. 启动成功

![启动项目](deploy/access/快速开始/9启动成功.jpg)

#### 

## 常见问题
### 1. go mod tidy 超时 i/o timeout
```
1. 确认当前shell
echo $SHELL

2. 编辑相应的 Shell 配置文件
a. 如果使用 zsh
vim ~/.zshrc

a. 如果使用 bash
vim ~/.bash_profile

3. 添加配置信息
export GOPROXY=https://proxy.golang.org,https://mirrors.aliyun.com/goproxy/,direct

4. 重载配置
source ~/.zshrc
source ~/.bash_profile

5. 查看更新结果
go env GOPROXY
```

## 许可证

本项目采用 Apache License 2.0 许可证 - 查看 LICENSE 文件了解详情
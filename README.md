### 一、概述

开箱即用的 go-zero 示例

内置服务api、scheduler、queue、script

符合Web应用框架的代码结构，便于开发者快速上手，减少开发成本。

### 二、技术选型

go-zreo 框架 + goctl 代码生成工具 + Makefile 实现自动化编译

go-zero 是一个集成了各种工程实践的 web 和 rpc 框架，缩短从需求到上线的距离，通过弹性设计保障了大并发服务端的稳定性，经受了充分的实战检验。

goctl 是 go-zero 微服务框架下的代码生成工具。使用 goctl 可显著提升开发效率，让开发人员将时间重点放在业务开发上，其功能有：api服务生成、rpc服务生成、model代码生成、模板管理

Makefile 文件描述了 Linux 系统下项目工程的编译规则，只需要一个 `make bild` 命令，整个工程就开始自动构建项目环境，不再需要手动执行大量的 `go build` 命令，Makefile 文件定义了一系列规则，指明了源文件的编译顺序、依赖关系、是否需要重新编译等，可以输入 `make help` 查看命令集。

### 三、代码结构

```text
.
├── app                             包含应用程序的主要代码
│   ├── doc                         项目主业务逻辑
│   ├── etc                         静态配置文件目录
│   ├── internal                    内部业务逻辑
│   ├── app.go                      应用程序的入口文件，定义了 API 服务、队列服务和调度服务的启动命令。
│   ├── wire.go                     依赖注入文件
│   └── wire_gin.go                 依赖注入生成文件
├── build                           项目构建目录
│   ├── app                         应用构建目录  
│   │   ├── bin                     应用二进制文件
│   │   └── deploy                  应用部署文件                 
│   │   └── etc                     配置文件         
│   └── app.tar                     应用构建打包文件（部署文件）         
├── deploy                          部署目录
│   ├── goctl                       代码生成工具                       
│   ├── sql                         数据库脚本
│   └── template                    模板文件 
├── pkg                             工具包
├── runtime                         项目运行时目录
├── vendor                          项目依赖包
├── .gitignore                      git 忽略文件
├── go.mod                          项目依赖管理文件
├── Makefile                        项目构建文件
└── README.md                       项目说明文件
```

### 四、环境要求
- Golang >= 1.20
- MySql >= 5.5.48

### 五、项目部署
```shell
# 生成api代码文件
$ make api

# 生成model代码文件
$ make model

# 生成wire代码文件
$ make wire

# 构建项目
$ make build env=dev|test|prod

```

### 六、安装
#### 1. go-zero
- 仓库地址：https://github.com/zeromicro/go-zero
- 文档地址：https://go-zero.dev/cn/

```shell
$ go get -u github.com/zeromicro/go-zero
```

#### 2. goctl
- 仓库地址：https://github.com/zeromicro/go-zero/tree/master/tools/goctl
- 文档地址：https://go-zero.dev/cn/docs/goctl/goctl

```shell
$ go get -u github.com/zeromicro/go-zero/tools/goctl
```

### 七、常见问题
#### 1 go mod tidy 超时 i/o timeout
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
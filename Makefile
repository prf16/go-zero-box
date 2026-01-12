export GO111MODULE=on
export GOPROXY=https://goproxy.cn
export GOSUMDB=sum.golang.org

.PHONY: build
# 构建并打包应用（根据 env=dev|test|prod 编译，生成 build/app 及 app.tar）
build:
	$(info ******************** build ********************)
	@echo "process [build] env=$(env)"

	@if [ "$(env)" = "dev" ]; then \
  		echo "Building for dev environment"; \
		go build -ldflags="-s -w" -o ./build/app/app app/app.go app/wire_gen.go;\
	elif [ "$(env)" = "prod" ]; then \
  		echo "Building for prod environment"; \
		GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./build/app/app app/app.go app/wire_gen.go;\
	elif [ "$(env)" = "test" ]; then \
		echo "Building for test environment"; \
		GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./build/app/app app/app.go app/wire_gen.go;\
	else \
		echo "Error: Unknown set env. Please set env to 'dev', 'prod', 'test'. Example：make api env=prod"; \
		exit 1; \
	fi

	mkdir -p ./build/app/etc
	mkdir -p ./build/app/api
	cp -f app/etc/app.yaml.$(env).bak ./build/app/etc/app.yaml
	cp -f app/api/api.json ./build/app/api/api.json
	tar -C ./build -cvf ./build/app.tar app

.PHONY: api
# 根据 api.api 定义生成 Go API 代码与 Swagger 文档
api:
	$(info ******************** api ********************)
	@echo "process build [api]"
	goctl api go -api app/api/api.api -dir app --style go_zero --home ./deploy/goctl/1.5.5/
	goctl api swagger --api app/api/api.api --dir app/api/
	@echo "processed"

.PHONY: wire
# 根据 wire.go 生成依赖注入代码（wire_gen.go）
wire:
	$(info ******************** wire ********************)
	@echo "process build [wire]"
	wire app/wire.go
	@echo "processed"

.PHONY: model
# 根据 MySQL 表结构生成 Go Model 代码
model:
	$(info ******************** model ********************)
	@echo "process model"
	goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/user" -table="user*" -dir app/internal/model/usermodel --style go_zero -i 'created_at,updated_at' --home ./deploy/goctl/1.5.5/
	goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/message" -table="message*" -dir app/internal/model/messagemodel --style go_zero -i 'created_at,updated_at' --home ./deploy/goctl/1.5.5/
	@echo "processed"

# 帮助
help:
	@echo 'Version: v1.9.4-0.0.4'
	@echo ''
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
    helpMessage = match(lastLine, /^# (.*)/); \
        if (helpMessage) { \
            helpCommand = substr($$1, 0, index($$1, ":")-1); \
            helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
            printf "    \033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
        } \
    } \
    { lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
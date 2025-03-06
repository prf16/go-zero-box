package config

import (
	"github.com/google/wire"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-box/pkg/asynqx"
)

var Provider = wire.NewSet(
	NewConfig,
	NewDBTest,
	NewDBTest1,
	NewRedis,
)

var globalConfig = new(Config)

type Config struct {
	Api       *Api
	DB        *DB
	JwtAuth   *JwtAuth
	Redis     *Redis
	Scheduler *asynqx.Config
	Queue     *asynqx.Config
}

func NewConfig() *Config {
	return globalConfig
}

func GetConfig() *Config {
	return globalConfig
}

// Api - api 配置
type Api struct {
	rest.RestConf
}

// JwtAuth - jwt 配置
type JwtAuth struct {
	AccessSecret string
	AccessExpire int64
}

func Init(c *Config) {
	globalConfig = c

	asynqClientInit()
}

func asynqClientInit() {
	err := asynqx.Init(GetConfig().Queue)
	if err != nil {
		logx.Must(err)
		return
	}
}

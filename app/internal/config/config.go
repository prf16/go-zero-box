package config

import (
	"github.com/google/wire"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-box/app/internal/pkg/database"
	"go-zero-box/app/internal/pkg/redis"
	"go-zero-box/pkg/asynqx"
)

var Provider = wire.NewSet(
	wire.FieldsOf(new(*Config), "Database", "Redis"),
)

type Config struct {
	Server    rest.RestConf
	Database  *database.Config
	Redis     *redis.Config
	Scheduler *asynqx.Config
	Queue     *asynqx.Config
	JwtAuth   *JwtAuth
}

// JwtAuth - jwt 配置
type JwtAuth struct {
	AccessSecret string
	AccessExpire int64
}

// todo 初始化位置修改
//func Init(c *Config) {
//	asynqClientInit()
//}
//
//func asynqClientInit() {
//	err := asynqx.Init(GetConfig().Queue)
//	if err != nil {
//		logx.Must(err)
//		return
//	}
//}

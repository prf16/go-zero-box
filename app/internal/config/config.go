package config

import (
	"go-zero-box/pkg/asynqx"
	"go-zero-box/pkg/database"
	"go-zero-box/pkg/redis"
	"go-zero-box/pkg/rpc"

	"github.com/google/wire"
	"github.com/zeromicro/go-zero/rest"
)

var Provider = wire.NewSet(
	wire.FieldsOf(new(*Config), "Database", "Redis", "Asynqx", "UserRpc"),
)

type Config struct {
	Server  rest.RestConf
	JwtAuth *JwtAuth

	Database *database.Config
	Redis    *redis.Config
	Asynqx   *asynqx.Config
	UserRpc  *rpc.Config
}

// JwtAuth - jwt 配置
type JwtAuth struct {
	AccessSecret string
	AccessExpire int64
}

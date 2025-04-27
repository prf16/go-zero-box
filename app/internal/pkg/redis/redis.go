package redis

import (
	"github.com/google/wire"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var Provider = wire.NewSet(
	NewDefault,
)

type Config struct {
	Host string
	Type string
	Pass string
}

type Default struct {
	*redis.Redis
}

func NewDefault(c *Config) *Default {
	return &Default{redis.MustNewRedis(redis.RedisConf{
		Host: c.Host,
		Type: c.Type,
		Pass: c.Pass,
	})}
}

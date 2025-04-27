package redis

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
)

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

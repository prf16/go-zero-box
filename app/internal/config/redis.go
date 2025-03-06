package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Redis struct {
	Host string
	Type string
	Pass string
}

func NewRedis(c *Config) *redis.Redis {
	return redis.MustNewRedis(redis.RedisConf{
		Host: c.Redis.Host,
		Type: c.Redis.Type,
		Pass: c.Redis.Pass,
	})
}

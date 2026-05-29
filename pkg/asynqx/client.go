package asynqx

import (
	"go-zero-box/pkg/redis"

	"github.com/hibiken/asynq"
)

var Client *asynq.Client

func NewClient(c *redis.Config) *asynq.Client {
	Client = asynq.NewClient(asynq.RedisClientOpt{
		Addr:     c.Host,
		Password: c.Pass,
	})
	return Client
}

func NewTask(typename string, payload []byte, opts ...asynq.Option) *asynq.Task {
	return asynq.NewTask(typename, payload, opts...)
}

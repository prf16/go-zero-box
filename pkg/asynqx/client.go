package asynqx

import (
	"github.com/hibiken/asynq"
)

var Client *asynq.Client

func NewClient(c *Config) *asynq.Client {
	Client = asynq.NewClient(asynq.RedisClientOpt{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})
	return Client
}

func NewTask(typename string, payload []byte, opts ...asynq.Option) *asynq.Task {
	return asynq.NewTask(typename, payload, opts...)
}

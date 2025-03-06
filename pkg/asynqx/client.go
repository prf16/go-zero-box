package asynqx

import (
	"github.com/hibiken/asynq"
)

var _client *asynq.Client

func Init(c *Config) error {
	_client = asynq.NewClient(asynq.RedisClientOpt{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})
	return nil
}
func GetClient() *asynq.Client {
	return _client
}

func NewTask(typename string, payload []byte, opts ...asynq.Option) *asynq.Task {
	return asynq.NewTask(typename, payload, opts...)
}

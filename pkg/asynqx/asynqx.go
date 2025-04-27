package asynqx

import (
	"github.com/google/wire"
	"github.com/hibiken/asynq"
)

var Provider = wire.NewSet(
	NewAsynq,
	NewClient,
)

type Asynq struct {
	Client *asynq.Client
}

func NewAsynq(client *asynq.Client) *Asynq {
	return &Asynq{Client: client}
}

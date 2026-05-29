package asynqx

import (
	"github.com/google/wire"
	"github.com/hibiken/asynq"
)

var Provider = wire.NewSet(
	NewAsynqx,
	NewClient,
)

type Asynqx struct {
	Client *asynq.Client
}

func NewAsynqx(client *asynq.Client) *Asynqx {
	return &Asynqx{Client: client}
}

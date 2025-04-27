package redis

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewRedis,
	NewDefault,
)

type Redis struct {
	Default *Default
}

func NewRedis(Default *Default) *Redis {
	return &Redis{Default: Default}
}

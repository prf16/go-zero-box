package rpc

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewRpc,
	NewUser,
)

type Rpc struct {
	User *User
}

func NewRpc(user *User) *Rpc {
	return &Rpc{User: user}
}

package rpc

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type User struct {
	zrpc.Client
}

func NewUser(c *Config) *User {
	return &User{zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: c.Target,
	})}
}

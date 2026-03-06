package rpc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type User struct {
	zrpc.Client
}

func NewUser(c *Config) *User {
	client, err := zrpc.NewClient(zrpc.RpcClientConf{
		Target: c.Target,
	})
	if err != nil {
		logx.Errorf("Warning: rpc.NewUser Fatal: %v", err)
	}

	return &User{client}
}

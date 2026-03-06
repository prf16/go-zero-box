package rpc

import (
	"github.com/prf16/go-zero-box-rpc/api/user"

	"github.com/zeromicro/go-zero/zrpc"
)

type User struct {
	user.UserClient
}

func NewUser(c *Config) *User {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: c.Target,
	}).Conn()

	return &User{user.NewUserClient(conn)}
}

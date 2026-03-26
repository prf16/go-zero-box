package rpc

import (
	"github.com/prf16/go-zero-box-rpc/api/user"
)

type User struct {
	user.UserClient
}

func NewUser(config *Config) *User {
	//client := zrpc.MustNewClient(zrpc.RpcClientConf{
	//	Target: config.Target,
	//})
	//
	//return &User{user.NewUserClient(client.Conn())}
	return &User{}
}

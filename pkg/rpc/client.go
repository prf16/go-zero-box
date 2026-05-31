package rpc

import (
	"github.com/prf16/go-zero-box-rpc/api/user"
)

type User struct {
	user.UserClient
}

func NewUser(config *Config) *User {
	//var i int
	//for {
	//	client, err := zrpc.NewClient(zrpc.RpcClientConf{
	//		Target: config.Target,
	//	})
	//	if err == nil {
	//		return &User{user.NewUserClient(client.Conn())}
	//	}
	//
	//	log.Printf("rpc.NewUser zrpc.NewClient failed retry:%d, err: %v", i, err)
	//	time.Sleep(2 * time.Second) // 两秒后重试
	//	i++
	//	if i > 3 {
	//		panic(err)
	//	}
	//}

	return &User{}
}

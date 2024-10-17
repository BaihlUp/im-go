package svc

import (
	"easy-chat/apps/user/api/internal/config"
	"easy-chat/apps/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
	// N * client =》 别名
)

type ServiceContext struct {
	Config config.Config

	userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		User: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}

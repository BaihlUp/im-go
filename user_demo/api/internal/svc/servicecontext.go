package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"im-go/api/internal/config"
	"im-go/api/internal/middleware"
	"im-go/rpc/userclient"
)

type ServiceContext struct {
	Config            config.Config
	LoginVerification rest.Middleware

	UserClient userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		LoginVerification: middleware.NewLoginVerificationMiddleware().Handle,
		UserClient:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}

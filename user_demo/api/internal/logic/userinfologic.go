package logic

import (
	"context"
	"im-go/rpc/userclient"

	"im-go/api/internal/svc"
	"im-go/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserinfoLogic) Userinfo(req *types.UserReq) (resp *types.UserResp, err error) {
	// todo: add your logic here and delete this line
	getUserResp, err := l.svcCtx.UserClient.GetUser(l.ctx, &userclient.GetUserReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserResp{
		Id:    getUserResp.Id,
		Name:  getUserResp.Name,
		Phone: getUserResp.Phone,
	}, nil
	return
}

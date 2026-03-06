package hello

import (
	"context"
	"go-zero-box/app/internal/utils/result"

	"go-zero-box/app/internal/svc"
	"go-zero-box/app/internal/types"

	"github.com/prf16/go-zero-box-rpc/api/user"
	"github.com/zeromicro/go-zero/core/logx"
)

type HelloUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloUserLogic {
	return &HelloUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloUserLogic) HelloUser(req *types.HelloUserReq) (resp *types.HelloUserResp, err error) {
	resp = &types.HelloUserResp{}

	userInfo, err := l.svcCtx.Pkg.Rpc.User.Info(l.ctx, &user.UserInfoReq{})
	if err != nil {
		return nil, result.ResponseSystem(l.ctx, err.Error())
	}

	resp.Data = types.HelloUserRespData{
		Id:       userInfo.User.ID,
		Nickname: userInfo.User.NickName,
	}

	return
}

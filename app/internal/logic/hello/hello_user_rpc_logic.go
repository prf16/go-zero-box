package hello

import (
	"context"
	"go-zero-box/app/internal/utils/result"
	"go-zero-box/app/rpc/user_rpc"

	"go-zero-box/app/internal/svc"
	"go-zero-box/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloUserRpcLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloUserRpcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloUserRpcLogic {
	return &HelloUserRpcLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloUserRpcLogic) HelloUserRpc(req *types.HelloUserRpcReq) (resp *types.HelloUserRpcResp, err error) {
	resp = &types.HelloUserRpcResp{}

	userInfo, err := user_rpc.NewUserClient(l.svcCtx.Pkg.Rpc.User.Conn()).Info(l.ctx, &user_rpc.UserInfoReq{})
	if err != nil {
		return nil, result.ResponseSystem(l.ctx, err.Error())
	}

	resp.Data = types.HelloUserRpcRespData{
		Id:       userInfo.User.ID,
		Nickname: userInfo.User.NickName,
	}
	return
}

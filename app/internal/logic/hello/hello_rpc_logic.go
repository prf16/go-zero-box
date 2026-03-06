package hello

import (
	"context"

	"go-zero-box/app/internal/svc"
	"go-zero-box/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloRpcLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloRpcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloRpcLogic {
	return &HelloRpcLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloRpcLogic) HelloRpc(req *types.HelloRpcReq) (resp *types.HelloRpcResp, err error) {
	// todo: add your logic here and delete this line

	return
}

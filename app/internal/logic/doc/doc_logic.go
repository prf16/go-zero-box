package doc

import (
	"context"

	"go-zero-box/app/internal/svc"
	"go-zero-box/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DocLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDocLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DocLogic {
	return &DocLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DocLogic) Doc(req *types.SwaggerReq) (resp *types.SwaggerResp, err error) {
	// todo: add your logic here and delete this line

	return
}

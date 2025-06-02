package doc

import (
	"context"

	"go-zero-box/app/internal/svc"
	"go-zero-box/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DocJsonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDocJsonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DocJsonLogic {
	return &DocJsonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DocJsonLogic) DocJson(req *types.SwaggerReq) (resp *types.SwaggerResp, err error) {
	// todo: add your logic here and delete this line

	return
}

package doc

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-box/app/internal/svc"
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

func (l *DocJsonLogic) DocJson() error {
	// todo: add your logic here and delete this line

	return nil
}

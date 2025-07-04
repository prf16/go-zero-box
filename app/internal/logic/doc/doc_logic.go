package doc

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-box/app/internal/svc"
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

func (l *DocLogic) Doc() error {
	// todo: add your logic here and delete this line

	return nil
}

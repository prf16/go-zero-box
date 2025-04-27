package info

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"go-zero-box/app/internal/result"
	"go-zero-box/app/internal/svc"
	"go-zero-box/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExitLogic {
	return &ExitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExitLogic) Exit() (resp *types.ExitResp, err error) {

	resp = &types.ExitResp{}

	// 校验用户信息
	userId := cast.ToInt64(l.ctx.Value("user_id"))
	if userId == 0 {
		return nil, result.Response(l.ctx, "无法获取用户信息")
	}

	// 删除用户登入标识
	_, err = l.svcCtx.Pkg.Redis.Default.Del(fmt.Sprintf("user:auth:%d", userId))
	if err != nil {
		return nil, result.ResponseSystem(l.ctx, err.Error())
	}

	return
}

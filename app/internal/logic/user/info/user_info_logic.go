package info

import (
	"context"
	"github.com/spf13/cast"
	"go-zero-box/app/internal/svc"
	"go-zero-box/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResp, err error) {
	userID := cast.ToInt64(l.ctx.Value("user_id"))

	user, err := l.svcCtx.Model.UserModel.FindOne(l.ctx, userID)
	if err != nil {
		return nil, err
	}

	// 用户基本信息
	data := &types.UserInfoData{
		Id:      user.Id,
		Account: user.Account,
		Name:    user.Name,
		Email:   user.Email,
		Note:    user.Note,
		Status:  user.Status,
	}

	return &types.UserInfoResp{
		Data: data,
	}, nil
}

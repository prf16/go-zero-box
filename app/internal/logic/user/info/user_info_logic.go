package info

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/spf13/cast"
	"go-zero-box/app/internal/svc"
	"go-zero-box/app/internal/types"
	"go-zero-box/app/internal/utils/result"

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
	userID := cast.ToUint64(l.ctx.Value("user_id"))

	user, err := l.svcCtx.Model.UserModel.First(l.ctx, squirrel.Select("*").Where("id =?", userID))
	if err != nil {
		return nil, result.ResponseSystem(l.ctx, err.Error())
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

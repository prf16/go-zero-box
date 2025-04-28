package auth

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"go-zero-box/app/internal/model/usermodel"
	"go-zero-box/app/internal/svc"
	"go-zero-box/app/internal/types"
	"go-zero-box/app/internal/types/constant"
	"go-zero-box/app/internal/types/result"
	"go-zero-box/app/internal/types/tools"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	resp = &types.RegisterResp{}

	user, err := l.svcCtx.Model.UserModel.First(l.ctx, squirrel.Select("*").Where("account = ?", req.Account))
	if err != nil {
		return nil, result.ResponseSystem(l.ctx, err.Error())
	}
	if user.Id > 0 {
		return nil, result.Response(l.ctx, "账号已存在")
	}
	// 2.添加用户
	_, err = l.svcCtx.Model.UserModel.Insert(l.ctx, &usermodel.User{
		Uuid:      uuid.New().String(),
		Account:   req.Account,
		Password:  tools.HashAndSalt([]byte(req.Password)),
		Status:    constant.AccountStatusEnable,
		IsDelete:  constant.IsNoDelete,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, result.ResponseSystem(l.ctx, err.Error())
	}
	return
}

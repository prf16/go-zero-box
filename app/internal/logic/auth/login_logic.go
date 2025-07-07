package auth

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/golang-jwt/jwt/v4"
	"go-zero-box/app/internal/model/usermodel"
	"go-zero-box/app/internal/queue/message"
	"go-zero-box/app/internal/utils/constant"
	"go-zero-box/app/internal/utils/result"
	"go-zero-box/app/internal/utils/tools"
	"strings"
	"time"

	"go-zero-box/app/internal/svc"
	"go-zero-box/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	resp = &types.LoginResp{
		Data: types.LoginData{},
	}

	// 检查用户
	rowBuilder := squirrel.Select("*").Where("is_delete = ?", constant.IsNoDelete)
	if strings.Index(req.Account, "@") == -1 {
		rowBuilder = rowBuilder.Where(
			squirrel.Or{
				squirrel.Expr("account = ?", req.Account),
				squirrel.Expr("mobile = ?", req.Account),
			})
	} else {
		rowBuilder = rowBuilder.Where("email = ?", req.Account)
	}
	userInfo, err := l.svcCtx.Model.UserModel.First(l.ctx, rowBuilder)
	if err != nil {
		return nil, result.ResponseSystem(l.ctx, err.Error())
	}

	// 校验用户是否存在
	if userInfo == nil || userInfo.Id == 0 {
		return nil, result.Response(l.ctx, result.MessageAuthNotUser)
	}
	if userInfo.Status != constant.AccountStatusEnable {
		return nil, result.Response(l.ctx, result.MessageAuthUserStatusDisable)
	}

	// 密码校验工作操作--密码校验
	if !tools.ValidatePasswords(userInfo.Password, []byte(req.Password)) {
		return nil, result.Response(l.ctx, "账号或密码不正确")
	}

	// 生成token
	jwtAuth, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessExpire, userInfo)
	if err != nil {
		return nil, result.ResponseSystem(l.ctx, err.Error())
	}

	err = message.MailQueueEnqueue(l.ctx, message.MailQueuePayload{
		User:    userInfo,
		Content: "登录成功",
	})
	if err != nil {
		return nil, result.ResponseSystem(l.ctx, err.Error())
	}

	resp.Data.Token = jwtAuth
	resp.Data.User = types.LoginDataUser{
		ID:       userInfo.Id,
		NickName: userInfo.Name,
	}

	return resp, nil
}

func (l *LoginLogic) getJwtToken(seconds int64, user *usermodel.User) (string, error) {
	iat := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["user_id"] = user.Id

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(l.svcCtx.Config.JwtAuth.AccessSecret))
}

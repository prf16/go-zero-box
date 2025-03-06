package middleware

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-box/app/internal/config"
	"go-zero-box/app/internal/result"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 用户授权
		authorization := r.Header.Get("Authorization")
		const bearerPrefix = "Bearer "
		if len(authorization) > len(bearerPrefix) && authorization[:len(bearerPrefix)] == bearerPrefix {
			authorization = authorization[len(bearerPrefix):]
		}

		token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetConfig().JwtAuth.AccessSecret), nil
		})

		if err != nil || !token.Valid {
			logx.ErrorStack(err.Error())
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, result.ResponseAuth(r.Context(), result.MessageAuthTokenNotValid))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, result.ResponseAuth(r.Context(), result.MessageAuthTokenNotValid))
			return
		}

		ctx := r.Context()
		for k, v := range claims {
			ctx = context.WithValue(ctx, k, v)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

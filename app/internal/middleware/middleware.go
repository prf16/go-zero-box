package middleware

import (
	"github.com/google/wire"
	"github.com/zeromicro/go-zero/rest"
)

var Provider = wire.NewSet(
	NewAuthMiddleware,
	NewLogMiddleware,
	NewMiddleware,
)

type Middleware struct {
	AuthMiddleware rest.Middleware
	LogMiddleware  rest.Middleware
}

func NewMiddleware(authMiddleware *AuthMiddleware, logMiddleware *LogMiddleware) *Middleware {
	return &Middleware{AuthMiddleware: authMiddleware.Handle, LogMiddleware: logMiddleware.Handle}
}

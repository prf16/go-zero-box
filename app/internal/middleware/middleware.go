package middleware

import (
	"github.com/google/wire"
	"github.com/zeromicro/go-zero/rest"
)

var Provider = wire.NewSet(
	NewAuthMiddleware,
	NewMiddleware,
)

type Middleware struct {
	AuthMiddleware rest.Middleware
}

func NewMiddleware(authMiddleware *AuthMiddleware) *Middleware {
	return &Middleware{AuthMiddleware: authMiddleware.Handle}
}

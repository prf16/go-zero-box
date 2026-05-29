package services

import (
	"go-zero-box/app/internal/svc/services/message"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewServices,
	message.NewService,
)

type Services struct {
	Message *message.Service
}

func NewServices(message *message.Service) *Services {
	return &Services{Message: message}
}

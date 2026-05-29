package services

import (
	"go-zero-box/app/internal/svc/services/demo"
	"go-zero-box/app/internal/svc/services/message"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewServices,
	demo.NewService,
	message.NewService,
)

type Services struct {
	Demo    *demo.Service
	Message *message.Service
}

func NewServices(demo *demo.Service, message *message.Service) *Services {
	return &Services{Demo: demo, Message: message}
}

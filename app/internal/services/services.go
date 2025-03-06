package services

import (
	"github.com/google/wire"
	"go-zero-box/app/internal/services/demo"
	"go-zero-box/app/internal/services/message"
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

package model

import (
	"go-zero-box/app/internal/svc/model/messagemodel"
	"go-zero-box/app/internal/svc/model/usermodel"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewModel,
	usermodel.NewUserModel,
	messagemodel.NewMessageModel,
)

type Model struct {
	UserModel    usermodel.UserModel
	MessageModel messagemodel.MessageModel
}

func NewModel(userModel usermodel.UserModel, messageModel messagemodel.MessageModel) *Model {
	return &Model{UserModel: userModel, MessageModel: messageModel}
}

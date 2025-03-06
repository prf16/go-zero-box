package model

import (
	"github.com/google/wire"
	"go-zero-box/app/internal/model/messagemodel"
	"go-zero-box/app/internal/model/usermodel"
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

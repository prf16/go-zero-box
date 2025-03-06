package demo

import (
	"go-zero-box/app/internal/model/usermodel"
)

type Service struct {
	UserModel usermodel.UserModel
}

func NewService(userModel usermodel.UserModel) *Service {
	return &Service{UserModel: userModel}
}

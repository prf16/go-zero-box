package usermodel

import (
	"go-zero-box/app/internal/config"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn config.DBTest) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn.SqlConn),
	}
}

// Init 初始化默认值
func (m *defaultUserModel) Init(data *User) {
}

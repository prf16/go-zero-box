package usermodel

import "go-zero-box/pkg/database"

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
func NewUserModel(conn *database.Default) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}

// Init 初始化默认值
func (m *defaultUserModel) Init(data *User) {
}

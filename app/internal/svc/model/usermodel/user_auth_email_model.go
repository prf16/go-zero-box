package usermodel

import "go-zero-box/pkg/database"

var _ UserAuthEmailModel = (*customUserAuthEmailModel)(nil)

type (
	// UserAuthEmailModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAuthEmailModel.
	UserAuthEmailModel interface {
		userAuthEmailModel
	}

	customUserAuthEmailModel struct {
		*defaultUserAuthEmailModel
	}
)

// NewUserAuthEmailModel returns a model for the database table.
func NewUserAuthEmailModel(conn *database.Default) UserAuthEmailModel {
	return &customUserAuthEmailModel{
		defaultUserAuthEmailModel: newUserAuthEmailModel(conn),
	}
}

// Init 初始化默认值
func (m *defaultUserAuthEmailModel) Init(data *UserAuthEmail) {
}

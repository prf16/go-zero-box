package messagemodel

import (
	"go-zero-box/pkg/database"
)

var _ MessageModel = (*customMessageModel)(nil)

type (
	// MessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageModel.
	MessageModel interface {
		messageModel
	}

	customMessageModel struct {
		*defaultMessageModel
	}
)

// NewMessageModel returns a model for the database table.
func NewMessageModel(conn *database.Default) MessageModel {
	return &customMessageModel{
		defaultMessageModel: newMessageModel(conn),
	}
}

// Init 初始化默认值
func (m *defaultMessageModel) Init(data *Message) {
}

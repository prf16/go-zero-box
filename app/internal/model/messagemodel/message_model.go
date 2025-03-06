package messagemodel

import (
	"go-zero-box/app/internal/config"
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
func NewMessageModel(conn config.DBTest1) MessageModel {
	return &customMessageModel{
		defaultMessageModel: newMessageModel(conn.SqlConn),
	}
}

// Init 初始化默认值
func (m *defaultMessageModel) Init(data *Message) {
}

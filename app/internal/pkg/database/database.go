package database

import (
	"github.com/google/wire"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var Provider = wire.NewSet(
	NewDefault,
)

// Config 数据库连接信息
type Config struct {
	Default string // 默认数据库
}

type Default struct {
	sqlx.SqlConn
}

func NewDefault(c *Config) *Default {
	return &Default{sqlx.NewMysql(c.Default)}
}

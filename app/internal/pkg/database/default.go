package database

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Default struct {
	sqlx.SqlConn
}

func NewDefault(c *Config) *Default {
	return &Default{sqlx.NewMysql(c.Default)}
}

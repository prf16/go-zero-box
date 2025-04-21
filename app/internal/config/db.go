package config

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DatabaseCore struct {
	SqlConn sqlx.SqlConn
}

func NewDatabaseCore(c *Config) DatabaseCore {
	return DatabaseCore{SqlConn: sqlx.NewMysql(c.Database.Core)}
}

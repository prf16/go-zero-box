package config

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// DB - 数据库
type DB struct {
	Test  string
	Test1 string
}

type DBTest struct {
	SqlConn sqlx.SqlConn
}
type DBTest1 struct {
	SqlConn sqlx.SqlConn
}

// NewDBTest Test数据库
func NewDBTest(c *Config) DBTest {
	return DBTest{SqlConn: sqlx.NewMysql(c.DB.Test)}
}

// NewDBTest1 Test1数据库
func NewDBTest1(c *Config) DBTest1 {
	return DBTest1{SqlConn: sqlx.NewMysql(c.DB.Test1)}
}

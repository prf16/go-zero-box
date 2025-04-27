package database

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewDatabase,
	NewDefault,
)

type Database struct {
	Default *Default
}

func NewDatabase(Default *Default) *Database {
	return &Database{Default: Default}
}

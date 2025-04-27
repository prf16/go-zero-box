package pkg

import (
	"github.com/google/wire"
	"go-zero-box/app/internal/pkg/database"
	"go-zero-box/app/internal/pkg/redis"
)

var Provider = wire.NewSet(
	NewPkg,
	database.Provider,
	redis.Provider,
)

type Pkg struct {
	Database *database.Default
	Redis    *redis.Default
}

func NewPkg(database *database.Default, redis *redis.Default) *Pkg {
	return &Pkg{Database: database, Redis: redis}
}

package app

import (
	"database/sql"
	"go-wire-example/internal/db"

	"github.com/go-redis/redis/v8"
)

type App struct { // 最终需要的对象
	Dao db.Dao
	Db  *sql.DB
	Rdb redis.UniversalClient
}

func NewApp(db *sql.DB, rdb redis.UniversalClient, dao db.Dao) *App {
	return &App{
		Db:  db,
		Rdb: rdb,
		Dao: dao,
	}
}

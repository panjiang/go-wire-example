package db

import (
	"database/sql"
	"log"

	"go-wire-example/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, NewDao, wire.Bind(new(Dao), new(*dao)))

func New(cfg *config.Config) (db *sql.DB, cleanup func(), err error) {
	log.Println("Connect db")
	db, err = sql.Open("mysql", cfg.Database.Dsn)
	if err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}

	cleanup = func() {
		log.Println("Close db")
		db.Close()
	}
	return db, cleanup, nil
}

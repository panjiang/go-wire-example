package app

import "database/sql"

type App struct { // 最终需要的对象
	DB *sql.DB
}

func NewApp(db *sql.DB) *App {
	return &App{DB: db}
}

package main

import (
	"log"
)

func main() {
	app, err := InitApp() // 使用wire生成的injector方法获取app对象
	if err != nil {
		log.Fatal(err)
	}
	var version string
	row := app.DB.QueryRow("SELECT VERSION()")
	if err := row.Scan(&version); err != nil {
		log.Fatal(err)
	}
	log.Println(version)
}

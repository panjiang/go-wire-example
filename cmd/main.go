package main

import (
	"go-wire-example/internal/entry"
	"log"
)

func main() {
	log.Println("App init...")
	app, cleanup, err := entry.InitApp() // 使用wire生成的injector方法获取app对象
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		log.Println("App close...")
		cleanup()
		log.Println("App closed")
	}()
	log.Println("App inited")
	version, err := app.Dao.Version() // 调用Dao接口方法
	if err != nil {
		log.Fatal(err)
	}
	log.Println("db version:", version)
}

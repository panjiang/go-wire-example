// +build wireinject

package main

import (
	"gowire/internal/app"
	"gowire/internal/config"
	"gowire/internal/db"

	"github.com/google/wire"
)

func InitApp() (*app.App, error) {
	panic(wire.Build(config.Provider, db.Provider, app.NewApp)) // 调用wire.Build方法传入所有的依赖对象以及构建最终对象的函数得到目标对象
}

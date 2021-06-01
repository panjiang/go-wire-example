// +build wireinject

package entry

import (
	"go-wire-example/config"
	"go-wire-example/internal/app"
	"go-wire-example/internal/db"
	"go-wire-example/internal/rdb"

	"github.com/google/wire"
)

func InitApp() (*app.App, func(), error) {
	panic(wire.Build(config.Provider, db.Provider, rdb.Provider, app.NewApp)) // 调用wire.Build方法传入所有的依赖对象以及构建最终对象的函数得到目标对象
}

// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package entry

import (
	"go-wire-example/config"
	"go-wire-example/internal/app"
	"go-wire-example/internal/db"
	"go-wire-example/internal/rdb"
)

// Injectors from wire.go:

func InitApp() (*app.App, func(), error) {
	configConfig, err := config.New()
	if err != nil {
		return nil, nil, err
	}
	sqlDB, cleanup, err := db.New(configConfig)
	if err != nil {
		return nil, nil, err
	}
	universalClient, cleanup2, err := rdb.New(configConfig)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	dao := db.NewDao(sqlDB)
	appApp := app.NewApp(sqlDB, universalClient, dao)
	return appApp, func() {
		cleanup2()
		cleanup()
	}, nil
}

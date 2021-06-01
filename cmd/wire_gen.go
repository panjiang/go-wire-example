// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"gowire/internal/app"
	"gowire/internal/config"
	"gowire/internal/db"
)

// Injectors from wire.go:

func InitApp() (*app.App, error) {
	configConfig, err := config.New()
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.New(configConfig)
	if err != nil {
		return nil, err
	}
	appApp := app.NewApp(sqlDB)
	return appApp, nil
}
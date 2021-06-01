package rdb

import (
	"context"
	"go-wire-example/config"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New)

func New(cfg *config.Config) (redis.UniversalClient, func(), error) {
	log.Println("Connect rdb")
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:        cfg.Redis.Addrs,
		WriteTimeout: time.Second * 3,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		log.Println("Close rdb")
		rdb.Close()

	}
	return rdb, cleanup, nil
}

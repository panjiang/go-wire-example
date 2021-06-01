package config

import (
	"encoding/json"
	"os"

	"github.com/google/wire"
)

var Provider = wire.NewSet(New) // 将New方法声明为Provider，表示New方法可以创建一个被别人依赖的对象,也就是Config对象

type Config struct {
	Database database `json:"database"`
	Redis    redis    `json:"redis"`
}

type database struct {
	Dsn string `json:"dsn"`
}

type redis struct {
	Addrs []string `json:"addrs"`
}

func New() (*Config, error) {
	fp, err := os.Open("config/config.json")
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	var cfg Config
	if err := json.NewDecoder(fp).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

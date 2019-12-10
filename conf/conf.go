package conf

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/encoder/toml"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/file"
)

type Config struct {
	Database DatabaseConfig `json:"database" toml:"database"`
}

var c config.Config
var conf Config

func Load() {
	c = config.NewConfig()
	enc := toml.NewEncoder()
	err := c.Load(file.NewSource(
		file.WithPath("./config.toml"),
		source.WithEncoder(enc)))
	if err != nil {
		panic(err)
	}
	err = c.Scan(&conf)
	if err != nil {
		panic(err)
	}
}

func GetDataBaseConfig() DatabaseConfig {
	return conf.Database
}

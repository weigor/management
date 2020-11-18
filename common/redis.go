package common

import (
	"flag"
	"github.com/BurntSushi/toml"
)

var (
	redisPath string
)

type RedisConfig struct {
	Addr     string
	Username string
	Password string
	DB       int
}

func init() {
	flag.StringVar(&redisPath, "redisconf", "./config/redis.toml", "")
}

func RedisInit() *RedisConfig {
	var conf RedisConfig
	_, err := toml.DecodeFile(redisPath, &conf)
	if err != nil {
		panic(err)
	}
	return &RedisConfig{
		Addr:     conf.Addr,
		Username: conf.Username,
		Password: conf.Password,
		DB:       conf.DB,
	}
}

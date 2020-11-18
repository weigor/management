package common

import (
	"context"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/go-redis/redis/v8"
)

var (
	redisPath string
)

type RedisConfig struct {
	Url string
	DB  int
}

func init() {
	flag.StringVar(&redisPath, "redisconf", "./config/redis.toml", "")
}

func RedisInit() *redis.Client {
	var conf RedisConfig
	_, err := toml.DecodeFile(redisPath, &conf)
	if err != nil {
		panic(err)
	}
	cli := redis.NewClient(&redis.Options{
		Addr: conf.Url,
		DB:   conf.DB,
	})
	if err := cli.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	return cli
}

func TestGetRedisCli() *redis.Client {
	flag.Set("conf", "../../config/redis.toml")
	flag.Parse()
	return RedisInit()
}

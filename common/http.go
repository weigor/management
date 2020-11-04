package common

import (
	"flag"
	"github.com/BurntSushi/toml"
)
var (
	httpPath string
)

type HttpConfig struct {
	Port int
}


func init() {
	flag.StringVar(&httpPath, "conf", "./config/http.toml", "")
}

func HttpInit() *HttpConfig {
	var conf HttpConfig
	_, err := toml.DecodeFile(httpPath, &conf)
	if err != nil {
		panic(err)
	}
	return &HttpConfig{
		Port: conf.Port,
	}
}

package main

import (
	"flag"
	"monitor-center/pkg/server"
)

func main(){
flag.()
	s := server.NewServer()
	if err := s.Start(); err != nil {
		return func(){}, err
	}
}

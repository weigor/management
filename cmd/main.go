package main

import (
	"context"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"management/log"
	"management/pkg/server"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	flag.Parse()
	log.Init()
	defer log.LoggerEnd()
	log.Logger.Info("", zap.String("log", "启动成功"))
	t := server.NewServer()
	if err := t.Start(); err != nil {
		log.Logger.Info("", zap.Error(err))
		panic(fmt.Sprintf("err: %s", err))
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Logger.Info(fmt.Sprintf("server get signal %s", s.String()))
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()
			if err := t.Shutdown(ctx); err != nil {
				log.Logger.Warn("shutdown gin server failed", zap.Error(err))
			}
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"management/log"
	"net/http"
)

type Server struct {
	ginEngine  *gin.Engine
	httpServer *http.Server
	handler    *Handler
}

func NewServer() *Server {
	server := &Server{
		handler: NewHandler(),
	}
	server.ginEngine = server.newRouter()

	server.httpServer = &http.Server{Addr: ":"+ getHttpPort(), Handler: server.ginEngine}

	log.Logger.Info("start", zap.Any("addr", server.httpServer.Addr))
	return server
}

func (s *Server) Start() error {
	if err := s.httpServer.ListenAndServe(); err != nil {
		log.Logger.Warn("start gin server failed", zap.Error(err))
		return err
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func getHttpPort () string {
	cfg := common.GetHttpConfig()
	if cfg != nil && cfg.Port != "" {
		return cfg.Port
	}

	return "3000"
}




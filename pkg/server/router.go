package server

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) newRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	router := e.Group("/api/v1/")
//	router.Use(s.handler.RsaDecrypt())
	{
		router.POST("user/login", s.handler.Login())
	}
	group := e.Group("/api/")
	group.Use(s.handler.TokenMiddleware())
	group.Use(s.handler.Cors())
	group.POST("user/create", s.handler.CreateUser())
	{
		//group.Use(JWTAuth())
		group.POST("user/list", s.handler.QueryUserList())
		group.POST("user/delete", s.handler.DeleteUser())
		group.POST("user/update", s.handler.UpdateUser())
		group.POST("user/batchUpdate", s.handler.BatchUpdateBatch())

	}
	return e
}

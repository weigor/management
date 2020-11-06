package server

import "github.com/gin-gonic/gin"

func (s *Server) newRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()

	group := e.Group("/api/")
	{
		group.POST("user/list", s.handler.QueryUserList())
		group.POST("user/delete", s.handler.DeleteUser())
		group.POST("user/update", s.handler.UpdateUser())
		group.POST("user/create", s.handler.CreateUser())
		group.POST("user/batchUpdate", s.handler.BatchUpdateBatch())

	}
	return e
}

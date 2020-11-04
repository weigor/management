package server

import "github.com/gin-gonic/gin"

func (s *Server) newRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()

	group := e.Group("/api/")
	{
		group.POST("user/list", s.handler.QueryMachineList())
		group.POST("user/delete", s.handler.DeleteMachine())
		group.POST("user/update", s.handler.UpdateMachine())
		group.POST("user/create", s.handler.CreateMachine())

	}
	return e
}

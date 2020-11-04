package server

import "github.com/gin-gonic/gin"

func (s *Server) newRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()

	aRoute := e.Group("/api/v1/")
	{


	}


	return e
}


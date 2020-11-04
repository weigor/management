package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type baseResp struct {
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
	ErrMsg  string      `json:"err_msg,omitempty"`
}

func HttpResponse400(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, baseResp{Success: false, ErrMsg: err.Error()})
}

func HttpResponse200(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, baseResp{Success: true, Data: data})
}

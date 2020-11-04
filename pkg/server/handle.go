package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"management/common"
	"management/log"
	"management/model"
	serverModel "management/pkg/server/model"
	"management/pkg/server/register"
)

type Handler struct {
	*register.Service
}

func NewHandler() *Handler {
	return &Handler{Service: register.NewService()}
}

func (h *Handler) QueryUserList() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (h *Handler) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &serverModel.UserReq{}
		if err := c.BindJSON(req); err != nil {
			log.Logger.Error("#CreateUser parse param fail", zap.Error(err))
			common.HttpResponse400(c, common.ParamsInvalidErr)
			return
		}
		if err := req.CreateVerification(); err != nil {
			log.Logger.Error("#CreateUser IsValid", zap.Error(err))
			common.HttpResponse400(c, common.ParamsInvalidErr)
			return
		}
		ctx := &serverModel.UserCtx{
			Req: &model.User{UserName: req.UserName,
				PassWord: req.PassWord,
				Age:      req.Age,
				Tel:      req.Tel,
				Addr:     req.Addr,
				Card:     req.Card,
			},
		}
		if err := h.UserService.Create(ctx); err != nil {
			log.Logger.Error("#CreateUser ", zap.Any("req", req), zap.Error(err))
			common.HttpResponse400(c, common.ParamsInvalidErr)
			return
		}
		common.HttpResponse200(c, nil)
	}
}

func (h *Handler) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (h *Handler) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

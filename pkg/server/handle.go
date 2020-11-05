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
		req := &serverModel.UserPageReq{}
		if err := c.BindJSON(req); err != nil {
			log.Logger.Error("#QueryUserList property resolution failed ", zap.Error(err))
			common.HttpResponse400(c, common.ParamsInvalidErr)
			return
		}
		if err := req.QueryVerification(); err != nil {
			log.Logger.Error("#QueryUserList  property validation error ", zap.Error(err))
			common.HttpResponse400(c, common.ParamsValidateErr)
			return
		}
		ctx := &serverModel.UserCtx{
			Req: &model.User{
				UserName: req.UserName,
				PassWord: req.PassWord,
				Age:      req.Age,
				Tel:      req.Tel,
				Addr:     req.Addr,
			},
		}
		if err := h.UserService.Query(ctx); err != nil {
			log.Logger.Error("#QueryUserList ", zap.Any("req", req), zap.Error(err))
			common.HttpResponse400(c, common.SystemErr)
			return
		}
		common.HttpResponse200(c, ctx.GetResult())
	}
}

func (h *Handler) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &serverModel.UserReq{}
		if err := c.BindJSON(req); err != nil {
			log.Logger.Error("#CreateUser property resolution failed", zap.Error(err))
			common.HttpResponse400(c, common.ParamsInvalidErr)
			return
		}
		if err := req.CreateVerification(); err != nil {
			log.Logger.Error("#CreateUser  property validation error", zap.Error(err))
			common.HttpResponse400(c, common.ParamsValidateErr)
			return
		}
		ctx := &serverModel.UserCtx{
			Req: &model.User{
				UserName: req.UserName,
				PassWord: req.PassWord,
				Age:      req.Age,
				Tel:      req.Tel,
				Addr:     req.Addr,
			},
		}
		if err := h.UserService.Create(ctx); err != nil {
			log.Logger.Error("#CreateUser ", zap.Any("req", req), zap.Error(err))
			common.HttpResponse400(c, common.SystemErr)
			return
		}
		common.HttpResponse200(c, nil)
	}
}

func (h *Handler) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &serverModel.UserReq{}
		if err := c.BindJSON(req); err != nil {
			log.Logger.Error("#UpdateUser property resolution failed", zap.Error(err))
			common.HttpResponse400(c, common.ParamsInvalidErr)
			return
		}
		if err := req.UpdateVerification(); err != nil {
			log.Logger.Error("#UpdateUser property validation error", zap.Error(err))
			common.HttpResponse400(c, common.ParamsValidateErr)
			return
		}
		ctx := &serverModel.UserCtx{
			Req: &model.User{
				UserName: req.UserName,
				PassWord: req.PassWord,
				Age:      req.Age,
				Tel:      req.Tel,
				Addr:     req.Addr,
				BaseModel: &model.BaseModel{
					ID: req.Id,
				},
			},
		}
		if err := h.UserService.Update(ctx); err != nil {
			log.Logger.Error("#CreateUser ", zap.Any("req", req), zap.Error(err))
			common.HttpResponse400(c, common.ParamsInvalidErr)
			return
		}
		common.HttpResponse200(c, nil)
	}
}

func (h *Handler) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &serverModel.UserReq{}
		if err := c.BindJSON(req); err != nil {
			log.Logger.Error("#DeleteUser property validation error", zap.Error(err))
			common.HttpResponse400(c, common.ParamsValidateErr)
			return
		}
		if err := req.DeleteVerification(); err != nil {
			log.Logger.Error("#DeleteUser  property validation error", zap.Error(err))
			common.HttpResponse400(c, common.ParamsInvalidErr)
			return
		}
		ctx := &serverModel.UserCtx{
			Req: &model.User{
				UserName: req.UserName,
				PassWord: req.PassWord,
				Age:      req.Age,
				Tel:      req.Tel,
				Addr:     req.Addr,
				BaseModel: &model.BaseModel{
					ID: req.Id,
				},
			},
		}
		if err := h.UserService.Delete(ctx); err != nil {
			log.Logger.Error("#DeleteUser ", zap.Any("req", req), zap.Error(err))
			common.HttpResponse400(c, common.SystemErr)
			return
		}
		common.HttpResponse200(c, nil)
	}
}

func (h *Handler) BatchUpdateBatch() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &serverModel.UsersReq{}
		users := make([]*serverModel.UserReq, 0)
		if err := c.BindJSON(&users); err != nil {
			log.Logger.Error("#BatchUpdateBatch property validation error", zap.Error(err))
			common.HttpResponse400(c, common.ParamsInvalidErr)
			return
		}
		req.UserReq = users
		if err := req.IsValid(); err != nil {
			log.Logger.Error("#BatchUpdateBatch property validation error", zap.Error(err))
			common.HttpResponse400(c, common.ParamsValidateErr)
			return
		}

		temps := make([]*model.User, len(req.UserReq))
		for i, v := range req.UserReq {
			temps[i] = &model.User{
				UserName: v.UserName,
				PassWord: v.PassWord,
				Age:      v.Age,
				Tel:      v.Tel,
				Addr:     v.Addr,
				BaseModel: &model.BaseModel{
					ID: v.Id,
				},
			}
		}
		ctx := &serverModel.UserCtx{
			Req: temps,
		}
		if err := h.UserService.BatchUpdate(ctx); err != nil {
			log.Logger.Error("#BatchUpdateBatch ", zap.Any("req", req), zap.Error(err))
			common.HttpResponse400(c, common.SystemErr)
			return
		}
		common.HttpResponse200(c, nil)
	}
}

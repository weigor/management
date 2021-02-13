package server

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/onsi/gomega/gbytes"
	"go.uber.org/zap"
	"io/ioutil"
	"management/common"
	"management/log"
	"management/model"
	serverModel "management/pkg/server/model"
	"management/pkg/server/register"
	"net/http"
)

type Handler struct {
	*register.Service
}

func NewHandler() *Handler {
	return &Handler{Service: register.NewService()}
}
func (h *Handler) RsaDecrypt() gin.HandlerFunc {
	return func(c *gin.Context) {
		body := c.Request.Body
		if body == nil {
			return
		}

		b, err := ioutil.ReadAll(body)
		if err != nil {
			log.Logger.Warn("RsaDecrypt", zap.Error(err))
			common.HttpResponse400(c, err)
			c.Abort()
			return
		}

		b, err = base64.StdEncoding.DecodeString(string(b))
		if err != nil {
			log.Logger.Warn("RsaDecrypt", zap.Error(err))
			common.HttpResponse400(c, err)
			c.Abort()
			return
		}

		//jsonb, err := h.rsaDecrypt(b)
		//if err != nil {
		//	log.Logger.Warn("RsaDecrypt", zap.Error(err))
		//	common.HttpResponse400(c, err)
		//	c.Abort()
		//	return
		//}

		c.Request.Body = gbytes.BufferWithBytes(b)
	}
}

func (h *Handler) rsaDecrypt(ciphertext []byte) ([]byte, error) {
	key := []byte("-----BEGIN RSA PRIVATE KEY-----\nMIIEpgIBAAKCAQEA0CXaTQE4ifUm2+tlH1oBLif6yWT7vRnnCqEvX8JUcmLQMrVu\n9WVMZbKDHphGMt6dFtQpCaXcrjPyPLqouel3lbro8vgIuPNSoeE1OZVbVWElvZ2V\npEvSA1rPe0bZ5F6A99Ts0Rz6uulVQkIkMQoxc4KFgrmBk+3a3VALsZLcEHSCsnnZ\nhnr6jYYSdh2DQvhIFqNaptFkEtb/xLLP9f0Xv2iVIq7wpMWByficDjzmuWxVeor+\nAUIy9fGcBepxQCgKq1yXhY/0zDUbgUUB/lxIGoQCDEKLRwJyQw/BbNeY5NFlZGqA\n3uPOoKhyMv5Pi/g5E4opSBx5W/lv3gKpB1cfpQIDAQABAoIBAQDGm2ep1EFjeXSj\noP8zJAk+Rk2IPv/pFr8aqGPwphc3sctgpzgBlK+J1gRAfCF3RmxzrOqfVxCzc8Nu\naNi30+oUB21g8IQ6HYo6Bg5oLHgihnihbaysQOBZ7RtOUHN18Spzz0pL2a/wCtYc\nS8oGtOgshFzqOCFIykrsowUVYcDzPM7E7KTaAoGCOhoL02Nw2iqtfHWkK/CuiSEG\nDoV4yLh6cSlSemKmN+kkcXICtbvAfLQFuKr1jUAC/3xW0MbnZtfkbbhb/0OBI1ri\nnYKgZbe5D9ne8nuZrcMH7D5FetoS5Ax7SIIZWYxmXGx74mxbeQRyfTLwLXdOGRGC\nup8VdHshAoGBAPpTH9q28SAgYKcQmCX+B3Aee/x1zmhfUPhYnGTcFQxMcfkUdTfx\noOJswWE+XRERkLxpRnIhpzgtJp2+pWhJsE6fFNMxML5fK5yeEF6J1EbCyGEwOtst\nvvigLs5oR9MlbF7nzopRMViN43e5BjkUD22yI5EMUZW4ubqa8mIlsN3NAoGBANTd\n73hywAubUnmlw4c2uoUzmINOxX+aUID87ONn0+fzY3Ug9LtJ6J0R88P4qY1wgdrU\nQjCMFgzmCH9q1Am8W1FqJkjGS/tWgVOIrUyfBRJO+EIlMIellR6n9DXCw9LGeO2V\nI7m6U1NWZGa8mtl/0S54QhNsfsp8p5twWV1tB7E5AoGBAKDzNXYRTnRTnRGOD+XN\nscabMykeLfrZ3lvvzY7kGvxvYpC+YKf5ynILb0MxL/G7k44xOkRD8xqhnUSrwfqN\n9rh2fJNV+3tMAeSPlQLUKBLfRquGsTEf9rwxcibw0c2nMEjNTvWMQugnQuxFoQSu\nK0Vi1o96ljJoNbMP0Wzdwxy5AoGBAK1kuxRaJKVPuDbvF/6kTfsCtFEBcU8n3Du1\nyyDSCoL+dx2J4tBMu/Z2ESKpAzP7WUtvaxswgSWwm2tvEZl8nMYMuXK+VFY/eMka\npE+tmOv497CpqoZUEswN85d3NxwSH58nxRoc9JMF5HLrXxecTkCUJP69eepm8ABl\n2+WGUqXBAoGBANoU+NjYsC5XVdk37iGjjLlFdHVX6G9AB1n2DSFGxS1qzJrwYk0C\n0kY+pDLzz9QLw+T7Sp6aochrgxZY97suGRnzw/jf4lVG48lg8d2dqXiqJGMmKLXh\n9vN4F6/II8r+aEDkHJU07VCTXvOnE6zJaSbsJiGWaKOqFTw3GnTQYiok\n-----END RSA PRIVATE KEY-----")
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes) // 解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func (h *Handler) TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		name := c.Request.Header.Get("name")
		err := h.UserService.Auth(name, token)
		if err == nil {
			c.Set("name", name)
			return
		}
		common.HttpResponse400(c, err)
		c.Abort()
	}
}
func (h *Handler) Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, token,name")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
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
		ctx := &serverModel.CommonCtx{
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
		common.HttpResponse200(c, ctx.GetResult(), http.StatusOK)
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
		ctx := &serverModel.CommonCtx{
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
		common.HttpResponse200(c, nil, http.StatusOK)
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
		ctx := &serverModel.CommonCtx{
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
		common.HttpResponse200(c, nil, http.StatusOK)
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
		ctx := &serverModel.CommonCtx{
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
		common.HttpResponse200(c, nil, http.StatusOK)
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
		ctx := &serverModel.CommonCtx{
			Req: temps,
		}
		if err := h.UserService.BatchUpdate(ctx); err != nil {
			log.Logger.Error("#BatchUpdateBatch ", zap.Any("req", req), zap.Error(err))
			common.HttpResponse400(c, common.SystemErr)
			return
		}
		common.HttpResponse200(c, nil, http.StatusOK)
	}
}

func (h *Handler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &serverModel.UserReq{}
		if err := c.BindJSON(req); err != nil {
			log.Logger.Error("#Login property resolution failed ", zap.Error(err))
			common.HttpResponse400(c, common.ParamsInvalidErr)
			return
		}
		if err := req.LoginVerification(); err != nil {
			log.Logger.Error("#Login  property validation error ", zap.Error(err))
			common.HttpResponse400(c, common.ParamsValidateErr)
			return
		}
		ctx := &serverModel.CommonCtx{
			Req: &model.User{
				UserName: req.UserName,
				PassWord: req.PassWord,
			},
		}
		if err := h.UserService.Login(ctx); err != nil {
			log.Logger.Error("#Login ", zap.Any("req", req), zap.Error(err))
			common.HttpResponse400(c, common.SystemErr)
			return
		}
		common.HttpResponse200(c, ctx.GetResult(), http.StatusOK)
	}
}

func (h *Handler) CreateLive() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &serverModel.LiveReq{}
		if err := c.BindJSON(req); err != nil {
			log.Logger.Error("#CreateLive property resolution failed ", zap.Error(err))
			common.HttpResponse400(c, common.ParamsInvalidErr)
			return
		}
		if err := req.CreateVerification(); err != nil {
			log.Logger.Error("#CreateLive  property validation error ", zap.Error(err))
			common.HttpResponse400(c, common.ParamsValidateErr)
			return
		}

		ctx := &serverModel.CommonCtx{
			Req: &model.Live{
				Head:     req.Head,
				Photo:    req.Photo,
				Username: req.Username,
				BaseModel: &model.BaseModel{
					Remark: req.Remark,
				},
			},
		}
		if err := h.LiveService.CreateLive(ctx); err != nil {
			log.Logger.Error("#CreateLive ", zap.Any("req", req), zap.Error(err))
			common.HttpResponse400(c, common.SystemErr)
			return
		}
		common.HttpResponse200(c, nil, http.StatusOK)
	}
}

func (h *Handler) QueryLiveList() gin.HandlerFunc {
	return func(c *gin.Context) {
		//req := &serverModel.LivePageReq{}
		//if err := c.BindJSON(req); err != nil {
		//	log.Logger.Error("#QueryLiveList property resolution failed ", zap.Error(err))
		//	common.HttpResponse400(c, common.ParamsInvalidErr)
		//	return
		//}
		//if err := req.QueryVerification(); err != nil {
		//	log.Logger.Error("#QueryLiveList  property validation error ", zap.Error(err))
		//	common.HttpResponse400(c, common.ParamsValidateErr)
		//	return
		//}
		ctx := &serverModel.CommonCtx{
			Req: &model.Live{

				BaseModel: &model.BaseModel{

				},
			},
		}
		if err := h.LiveService.QueryLiveList(ctx); err != nil {
			log.Logger.Error("#QueryUserList ", zap.Any("req", ctx), zap.Error(err))
			common.HttpResponse400(c, common.SystemErr)
			return
		}
		common.HttpResponse200(c, ctx.GetResult(), http.StatusOK)
	}
}

func (h *Handler) UpdateLive() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &serverModel.LiveReq{}
		if err := c.BindJSON(req); err != nil {
			log.Logger.Error("#UpdateLive property resolution failed", zap.Error(err))
			common.HttpResponse400(c, common.ParamsInvalidErr)
			return
		}
		if err := req.UpdateVerification(); err != nil {
			log.Logger.Error("#UpdateLive property validation error", zap.Error(err))
			common.HttpResponse400(c, common.ParamsValidateErr)
			return
		}
		ctx := &serverModel.CommonCtx{
			Req: &model.Live{
				Head:     req.Head,
				Photo:    req.Photo,
				Username: req.Username,
				BaseModel: &model.BaseModel{
					Remark: req.Remark,
					ID:     req.Id,
				},
			},
		}
		if err := h.LiveService.UpdateLive(ctx); err != nil {
			log.Logger.Error("#UpdateLive ", zap.Any("req", req), zap.Error(err))
			common.HttpResponse400(c, common.ParamsInvalidErr)
			return
		}
		common.HttpResponse200(c, nil, http.StatusOK)
	}
}

func (h *Handler) DeleteLive() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &serverModel.LiveReq{}
		if err := c.BindJSON(req); err != nil {
			log.Logger.Error("#DeleteLive property validation error", zap.Error(err))
			common.HttpResponse400(c, common.ParamsValidateErr)
			return
		}
		if err := req.DeleteVerification(); err != nil {
			log.Logger.Error("#DeleteLive  property validation error", zap.Error(err))
			common.HttpResponse400(c, common.ParamsInvalidErr)
			return
		}
		ctx := &serverModel.CommonCtx{
			Req: &model.Live{

				BaseModel: &model.BaseModel{
					Remark: req.Remark,
					ID:     req.Id,
				},
			},
		}
		if err := h.LiveService.DeleteLive(ctx); err != nil {
			log.Logger.Error("#DeleteLive ", zap.Any("req", req), zap.Error(err))
			common.HttpResponse400(c, common.SystemErr)
			return
		}
		common.HttpResponse200(c, nil, http.StatusOK)
	}
}

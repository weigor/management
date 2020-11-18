package register

import (
	"management/common"
	dao "management/pkg/dao"
	service "management/pkg/service"
)

type Service struct {
	UserService *service.UserService
	LiveService *service.LiveService
}

func NewService() *Service {
	s := &Service{}
	buildService(s)
	return s
}

func buildService(s *Service) {
	orm := common.MysqlInit()
	userDao := dao.NewUserDAO(orm)
	liveDao := dao.NewLiveDAO(orm)
	jwt := common.NewJwt()
	userService := service.NewUserService(userDao, jwt)
	liveService := service.NewLiveService(liveDao)
	s.UserService = userService
	s.LiveService = liveService
}

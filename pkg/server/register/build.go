package register

import (
	"management/common"
	dao "management/pkg/dao"
	service "management/pkg/service"
)

type Service struct {
	UserService *service.UserService
}

func NewService() *Service {
	s := &Service{}
	buildService(s)
	return s
}

func buildService(s *Service) {
	orm := common.MysqlInit()
	userDao := dao.NewUserDAO(orm)
	userService := service.NewUserService(userDao)
	s.UserService = userService
}

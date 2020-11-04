package register

import "management/pkg/service"

type Service struct {
	UserService *service.UserService
}

func NewService() *Service {
	s := &Service{}
	buildService(s)
	return s
}

func buildService(s *Service) {

}

package register

type Service struct {

}

func NewService() *Service {
	s := &Service{}
	buildMysql(s)
	return s
}

func buildMysql(s *Service) {

}

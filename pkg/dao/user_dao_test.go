package dao

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"management/common"
	"management/model"
	"strconv"
	"testing"
)

type UserDAOTestSuite struct {
	suite.Suite
	*require.Assertions
	dao *UserDAO
}

func (s *UserDAOTestSuite) SetT(t *testing.T) {
	s.Suite.SetT(t)
	s.Assertions = require.New(t)
}

func (s *UserDAOTestSuite) SetupSuite() { // 在第一个test运行之前执行
	s.dao = NewUserDAO(common.TestMysqlInit())
}

func (s *UserDAOTestSuite) TearDownAllSuite() { // 执行完所有test之后执行

}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserDAOTestSuite))
}

func (s *UserDAOTestSuite) SetupTest() { // 每个test方法之前执行

}

func (s *UserDAOTestSuite) TestCreate() {
	s.NoError(s.dao.CreateUser(&model.User{
		UserName: "zhangliang66",
		PassWord: "123",
		Age:      18,
		Tel:      "13555",
		Addr:     "dddd",
	}))
}

func (s *UserDAOTestSuite) TestQuery() {
	temp := &model.User{}
	users, err := s.dao.QueryUserList(temp, 0, 0)
	s.NoError(err)
	for _, v := range users {
		fmt.Printf("%v", v)
	}
}

func (s *UserDAOTestSuite) TestUpdate() {
	s.NoError(s.dao.UpdateUser(&model.User{
		UserName: "zhangliang",
		PassWord: "123",
		Age:      19,
		Tel:      "13555",
		Addr:     "dddd",
		BaseModel: &model.BaseModel{
			ID: 3,
		},
	}))
}

func (s *UserDAOTestSuite) TestDelete() {
	temp := &model.User{
		BaseModel: &model.BaseModel{
			ID: 1,
		},
	}
	s.NoError(s.dao.DeleteUser(temp.ID))
}

func (s *UserDAOTestSuite) TestBatchUpdate() {

	users := make([]*model.User, 0)
	for i := 1; i < 5; i++ {
		users=append(users,&model.User{
			UserName: "zhangliang"+strconv.Itoa(i),
			PassWord: "123",
			Age:      18,
			Tel:      "55555",
			Addr:     "sss",
			BaseModel: &model.BaseModel{
				ID: uint(i),
			},
		})
	}
	s.NoError(s.dao.BatchUpdateUsers(users))
}

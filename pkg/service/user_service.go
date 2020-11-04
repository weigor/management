package service

import "management/model"

type UserDAO interface {
	CreateUser(user []*model.User) error
	QueryUserList(user []*model.User) ([]*model.User, error)
}

type UserService struct {
}

func (n *UserService) Create(ctx UserCtx) error {
	return nil
}

package service

import "management/model"

type UserDAO interface {
	CreateUser(user *model.User) error
	QueryUserList(user *model.User, pageNo, pageSize int) ([]*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
	BatchUpdateUsers(users []*model.User) error
}

type UserService struct {
	UserDao UserDAO
}

func NewUserService(dao UserDAO) *UserService {
	return &UserService{UserDao: dao}
}

func (n *UserService) Create(ctx UserCtx) error {
	req := ctx.Param().(*model.User)
	return n.UserDao.CreateUser(req)
}

func (n *UserService) Query(ctx UserCtx) error {
	req := ctx.Param().(*model.User)
	users, err := n.UserDao.QueryUserList(req, ctx.GetPage(), ctx.GetPageSize())
	if err != nil {
		return err
	}
	ctx.SetResult(users)
	return err
}

func (n *UserService) Update(ctx UserCtx) error {
	req := ctx.Param().(*model.User)
	return n.UserDao.UpdateUser(req)
}

func (n *UserService) Delete(ctx UserCtx) error {
	return n.UserDao.DeleteUser(ctx.Param().(*model.User).ID)
}

func (n *UserService) BatchUpdate(ctx UserCtx) error {
	return n.UserDao.BatchUpdateUser(ctx.Param().([]*model.User))
}

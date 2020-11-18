package service

import (
	"management/common"
	"management/model"
	"time"
)

type UserDAO interface {
	CreateUser(user *model.User) error
	QueryUserList(user *model.User, pageNo, pageSize int) ([]*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
	BatchUpdateUsers(users []*model.User) error
	JwtCache(username, token string, time time.Duration) error
	QueryUser(name string) (*model.User, error)
	GetToken(username string) (string, error)
	UpdateUserToken(user map[string]interface{}) error
}

type UserService struct {
	userDao UserDAO
	jwt     *common.Jwt
}

func NewUserService(dao UserDAO, jwt *common.Jwt) *UserService {
	return &UserService{userDao: dao, jwt: jwt}
}

func (n *UserService) Create(ctx CommonCtx) error {
	req := ctx.Param().(*model.User)
	//生成token
	token, err := n.jwt.CreateToken(req.UserName)
	if err != nil {
		return err
	}
	req.Token = token
	req.ExpireTime = time.Now().Add(time.Hour * 24 * 30).Unix()
	return n.userDao.CreateUser(req)
}

func (n *UserService) Auth(name, token string) error {
	// step1 解析token
	tokenName, err := n.jwt.DecodeToken(token)
	if err != nil {
		return err
	}

	if tokenName != name {
		return common.TokenErr
	}

	user, err := n.userDao.QueryUser(name)
	if err != nil {
		return err
	}

	// step2 校验token过期时间
	if time.Now().Unix() > user.ExpireTime {
		return common.TokenTimeErr
	}
	return nil
}

func (n *UserService) Login(ctx CommonCtx) error {
	req := ctx.Param().(*model.User)
	user, err := n.userDao.QueryUser(req.UserName)
	if err != nil {
		return err
	}
	if user.PassWord != req.PassWord {
		return common.PassWordErr
	}
	//生成token
	token, err := n.jwt.CreateToken(req.UserName)

	//更新token
	err = n.userDao.UpdateUserToken(map[string]interface{}{"user_name": user.UserName, "token": token})
	ctx.SetResult(token)
	return err
}

func (n *UserService) Query(ctx CommonCtx) error {
	req := ctx.Param().(*model.User)
	users, err := n.userDao.QueryUserList(req, ctx.GetPage(), ctx.GetPageSize())
	if err != nil {
		return err
	}
	ctx.SetResult(users)
	return err
}

func (n *UserService) Update(ctx CommonCtx) error {
	req := ctx.Param().(*model.User)
	return n.userDao.UpdateUser(req)
}

func (n *UserService) Delete(ctx CommonCtx) error {
	return n.userDao.DeleteUser(ctx.Param().(*model.User).ID)
}

func (n *UserService) BatchUpdate(ctx CommonCtx) error {
	return n.userDao.BatchUpdateUsers(ctx.Param().([]*model.User))
}

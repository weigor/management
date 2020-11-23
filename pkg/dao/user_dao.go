package dao

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"management/common"
	"management/model"
	"time"
)

type UserDAO struct {
	db *common.Orm
	r  *redis.Client
}

func NewUserDAO(db *common.Orm) *UserDAO {
	return &UserDAO{db: db}
}

func (dao *UserDAO) CreateUser(user *model.User) error {
	temp := &model.User{}
	err := dao.db.Where("user_name=?", user.UserName).Last(&temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return dao.db.Create(user).Error
}
func (dao *UserDAO) QueryUserList(user *model.User, pageNo, pageSize int) ([]*model.User, error) {
	var users []*model.User
	if pageSize <= 0 {
		pageSize = 10
	}
	if pageNo == 0 {
		err := dao.db.Where(user).Find(&users).Error
		if err != nil {
			return nil, err
		}
	}
	return users, dao.db.Where(user).Find(&users).Limit(pageSize).Offset((pageNo - 1) * pageSize).Error

}

func (dao *UserDAO) UpdateUser(user *model.User) error {
	temp := &model.User{}
	err := dao.db.Where("user_name=?", user.UserName).Last(&temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if temp.ID != user.ID && temp.ID != 0 {
		return common.UpdateErr
	}
	return dao.db.Model(&model.User{}).Where("id=?", user.ID).Updates(user).Error
}

func (dao *UserDAO) DeleteUser(id uint) error {
	return dao.db.Where("id=?", id).Delete(&model.User{}).Error
}

func (dao *UserDAO) BatchUpdateUsers(users []*model.User) error {
	db := dao.db.Begin().Model(&model.User{})
	defer db.Rollback()
	for _, v := range users {
		if err := db.Where("id=?", v.ID).Update(v).Error; err != nil {
			return err
		}
	}
	db.Commit()
	return nil
}

func (dao *UserDAO) JwtCache(username, token string, time time.Duration) error {
	return dao.r.Set(context.Background(), username, token, time).Err()
}

func (dao *UserDAO) GetToken(username string) (token string, err error) {
	return dao.r.Get(context.Background(), username).Result()
}

func (dao *UserDAO) QueryUser(name string) (*model.User, error) {
	temp := &model.User{}
	err := dao.db.Where("user_name =?",name).Last(&temp).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("帐号不存在")
	}
	if err != nil {
		return nil, err
	}
	return temp, err
}

func (dao *UserDAO) UpdateUserToken(user map[string]interface{}) error {
	temp := &model.User{}
	err := dao.db.Where("user_name=?", user["user_name"].(string)).Last(&temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return dao.db.Model(&model.User{}).Where("user_name=?", user["user_name"].(string)).Update(user).Error
}

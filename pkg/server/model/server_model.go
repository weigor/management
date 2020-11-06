package model

import "errors"

type UserReq struct {
	UserName string `json:"username" name:"用户名"`
	PassWord string `g json:"password" name:"密码"`
	Age      uint   `json:"age" name:"年龄"`
	Tel      string `json:"tel" name:"手机号"`
	Addr     string `json:"addr" name:"地址"`
	Id       uint   `json:"id"`
}

func (r *UserReq) CreateVerification() error {
	switch {
	case r.UserName == "":
		return errors.New("username is isValid")
	case r.PassWord == "":
		return errors.New("password is isValid")
	case r.Tel == "":
		return errors.New("tel is isValid")
	case r.Addr == "":
		return errors.New("addr is isValid")
	default:
		return nil
	}
}

func (r *UserReq) UpdateVerification() error {
	switch {
	case r.Id <= 0:
		return errors.New("id is isValid")
	case r.UserName == "":
		return errors.New("username is isValid")
	case r.PassWord == "":
		return errors.New("password is isValid")
	case r.Tel == "":
		return errors.New("tel is isValid")
	case r.Addr == "":
		return errors.New("addr is isValid")
	default:
		return nil
	}
}

func (r *UserReq) DeleteVerification() error {
	switch {
	case r.Id <= 0:
		return errors.New("id is isValid")
	default:
		return nil
	}
}

type UsersReq struct {
	UserReq []*UserReq
}

func (r *UsersReq) IsValid() error {
	for _, g := range r.UserReq {
		if err := g.UpdateVerification(); err != nil {
			return err
		}
	}
	return nil
}

type UserPageReq struct {
	UserName string `json:"username" name:"用户名"`
	PassWord string `g json:"password" name:"密码"`
	Age      uint   `json:"age" name:"年龄"`
	Tel      string `json:"tel" name:"手机号"`
	Addr     string `json:"addr" name:"地址"`
	Card     string `json:"card" name:"身份证"`
	Page     int    `json:"page_no"`
}

func (r *UserPageReq) QueryVerification() error {
	switch {
	case r.Page < 0:
		return errors.New("page_no is isValid")
	default:
		return nil
	}
}

package model

import "errors"

type UserReq struct {
	UserName string `json:"username" name:"用户名"`
	PassWord string `g json:"password" name:"密码"`
	Age      uint   `json:"age" name:"年龄"`
	Tel      string `json:"tel" name:"手机号"`
	Addr     string `json:"addr" name:"地址"`
	Card     string `json:"card" name:"身份证"`
}

func (r *UserReq) CreateVerification() error {
	switch {
	case r.Card == "":
		return errors.New("card is isValid")
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

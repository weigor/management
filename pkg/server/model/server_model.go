package model

import "errors"

type UserReq struct {
	UserName string `json:"username" name:"用户名"`
	PassWord string ` json:"password" name:"密码"`
	Age      uint   `json:"age" name:"年龄"`
	Tel      string `json:"tel" name:"手机号"`
	Addr     string `json:"addr" name:"地址"`
	Id       uint   `json:"id"`
}

func (r *UserReq) LoginVerification() error {
	switch {
	case r.UserName == "":
		return errors.New("username is isValid")
	case r.PassWord == "":
		return errors.New("password is isValid")
	default:
		return nil
	}
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
	PassWord string `json:"password" name:"密码"`
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

type LiveReq struct {
	Head     string `json:"head"`
	Remark   string `json:"remark"`
	Photo    string `json:"photo"`
	Username string `json:"username"`
	Id       uint   `json:"id"`
}

func (r *LiveReq) CreateVerification() error {
	switch {
	case r.Head == "":
		return errors.New("head is isValid")
	case r.Username == "":
		return errors.New("username is isValid")
	default:
		return nil
	}
}

func (r *LiveReq) UpdateVerification() error {
	switch {
	case r.Head == "":
		return errors.New("head is isValid")
	case r.Username == "":
		return errors.New("username is isValid")
	case r.Id <= 0:
		return errors.New("id is isValid")
	default:
		return nil
	}
}
func (r *LiveReq) DeleteVerification() error {
	switch {
	case r.Id <= 0:
		return errors.New("id is isValid")

	default:
		return nil
	}
}

type LivePageReq struct {
	Head     string `json:"head"`
	Remark   string `json:"remark"`
	Photo    string `json:"photo"`
	Username string `json:"username"`
	Page     int    `json:"page_no"`
}

func (r *LivePageReq) QueryVerification() error {
	switch {
	case r.Page < 0:
		return errors.New("page is isValid")
	default:
		return nil
	}
}

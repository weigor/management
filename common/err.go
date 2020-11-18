package common

import "errors"

var (
	ParamsValidateErr = errors.New("property validation error")
	ParamsInvalidErr  = errors.New("property resolution failed")
	SystemErr  = errors.New("system error")
	UpdateErr         = errors.New("姓名已被注册，请重新查看")
	PassWordErr =errors.New("密码错误")
	TokenErr =errors.New("token错误")
	TokenTimeErr =errors.New("token失效请重新登录")
)

package common

import "errors"

var (
	ParamsValidateErr = errors.New("property validation error")
	ParamsInvalidErr  = errors.New("property resolution failed")
	SystemErr  = errors.New("system error")
	UpdateErr         = errors.New("姓名已被注册，请重新查看")
)

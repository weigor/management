package common

import "errors"

var (
	ParamsValidateErr = errors.New("property validation error")
	ParamsInvalidErr  = errors.New("property resolution failed")
	SystemErr  = errors.New("system error")
	UpdateErr         = errors.New("姓名或身份证已被注册，请重新查看")
)

package common

import "errors"

var (
	ParamsParseErr    = errors.New("params parse error")
	ParamsValidateErr = errors.New("params validate error")
	ParamsInvalidErr  = errors.New("param invalid")

)


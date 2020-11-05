package service

type UserCtx interface {
	Param() interface{}
	GetResult() interface{}
	SetResult(interface{})
	GetPage()int
	GetPageSize()int
}


package service

type CommonCtx interface {
	Param() interface{}
	GetResult() interface{}
	SetResult(interface{})
	GetPage()int
	GetPageSize()int
}


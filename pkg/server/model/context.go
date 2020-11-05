package model

type UserCtx struct {
	Req      interface{}
	Result   interface{}
	Page     int
	PageSize int
}

func (q *UserCtx) Param() interface{} {
	return q.Req
}

func (q *UserCtx) GetResult() interface{} {
	return q.Result
}

func (q *UserCtx) SetResult(t interface{}) {
	q.Result = t
}
func (q *UserCtx) GetPage() int {
	return q.Page
}

func (q *UserCtx) GetPageSize() int {
	return q.PageSize
}

package model

type CommonCtx struct {
	Req      interface{}
	Result   interface{}
	Page     int
	PageSize int
}

func (q *CommonCtx) Param() interface{} {
	return q.Req
}

func (q *CommonCtx) GetResult() interface{} {
	return q.Result
}

func (q *CommonCtx) SetResult(t interface{}) {
	q.Result = t
}
func (q *CommonCtx) GetPage() int {
	return q.Page
}

func (q *CommonCtx) GetPageSize() int {
	return q.PageSize
}

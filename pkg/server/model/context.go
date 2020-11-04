package model

type UserCtx struct {
	Req    interface{}
	Result interface{}
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

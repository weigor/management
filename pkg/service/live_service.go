package service

import "management/model"

type LiveDao interface {
	CreateLive(live *model.Live) error
	QueryLiveList(live *model.Live, pageNo, pageSize int) ([]*model.Live, error)
	UpdateLive(live *model.Live) error
	DeleteLive(id uint) error
}

type LiveService struct {
	dao LiveDao
}

func NewLiveService(LiveDao LiveDao) *LiveService {
	return &LiveService{dao: LiveDao}
}

func (s LiveService) CreateLive(ctx CommonCtx) error {
	req := ctx.Param().(*model.Live)
	return s.dao.CreateLive(req)
}

func (s LiveService) QueryLiveList(ctx CommonCtx) error {
	req := ctx.Param().(*model.Live)
	lives, err := s.dao.QueryLiveList(req, ctx.GetPage(), ctx.GetPageSize())
	if err != nil {
		return err
	}
	ctx.SetResult(lives)
	return err

}

func (s LiveService) UpdateLive(ctx CommonCtx) error {
	req := ctx.Param().(*model.Live)
	return s.dao.UpdateLive(req)
}

func (s LiveService) DeleteLive(ctx CommonCtx) error {
	req := ctx.Param().(*model.Live)
	return s.dao.DeleteLive(req.ID)
}

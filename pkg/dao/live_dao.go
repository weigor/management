package dao

import (
	"github.com/go-redis/redis/v8"
	"management/common"
	"management/model"
)

type LiveDAO struct {
	db *common.Orm
	r  *redis.Client
}

func NewLiveDAO(db *common.Orm) *LiveDAO {
	return &LiveDAO{db: db}
}

func (dao *LiveDAO) CreateLive(live *model.Live) error {
	return dao.db.Create(live).Error
}
func (dao *LiveDAO) QueryLiveList(live *model.Live, pageNo, pageSize int) ([]*model.Live, error) {
	var lives []*model.Live
	if pageSize <= 0 {
		pageSize = 10000
	}
	if pageNo == 0 {
		err := dao.db.Where(live).Order("created_at").Find(&lives).Error
		if err != nil {
			return nil, err
		}
	}
	return lives, dao.db.Where(live).Order("created_at").Find(&lives).Limit(pageSize).Offset((pageNo - 1) * pageSize).Error

}

func (dao *LiveDAO) UpdateLive(live *model.Live) error {
	return dao.db.Model(&model.Live{}).Where("id=?", live.ID).Updates(live).Error
}

func (dao *LiveDAO) DeleteLive(id uint) error {
	return dao.db.Where("id=?", id).Delete(&model.Live{}).Error
}

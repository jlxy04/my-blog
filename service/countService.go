package service

import (
	"github.com/my-blog/dao"
	"github.com/my-blog/datasource"
	"github.com/my-blog/extend"
	"github.com/my-blog/models"
	"strconv"
)

type CountAllDataService interface {
	scheduleCountAll()
}

type CountKey int16

const (
	COUNTKEY_READ CountKey = 0
	COUNTKEY_LIKE CountKey = 1
)

func (k CountKey) String() string {
	switch k {
	case COUNTKEY_READ:
		return "READ"
	case COUNTKEY_LIKE:
		return "LIKE"
	default:
		return ""
	}
}

func NewCountAllDataService() countAllDataService {
	return countAllDataService{
		extendDao: dao.NewExtendDao(datasource.InstanceMaster()),
		countDao:  dao.NewCountDao(datasource.InstanceMaster()),
	}
}

type countAllDataService struct {
	extendDao *dao.ExtendDao
	countDao  *dao.CountDao
}

func (s countAllDataService) scheduleCountAll() {
	count := s.extendDao.CountAll()

	counts := make([]models.BlogCount, 0)

	counts = append(counts, models.BlogCount{
		Id:  extend.GenerateId(),
		Key: COUNTKEY_READ.String(),
		Val: strconv.Itoa(count.ReadCount),
	})

	counts = append(counts, models.BlogCount{
		Id:  extend.GenerateId(),
		Key: COUNTKEY_LIKE.String(),
		Val: strconv.Itoa(count.LikeCount),
	})

	s.countDao.UpdateData(counts)
}

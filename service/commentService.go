package service

import (
	"my-blog/dao"
	"my-blog/datasource"
	"my-blog/models"
)

type CommonService interface {

	//根据文章ID查询评论列表，默认只显示最近10条评论
	ListCommonByArticleId(articleId string) []models.BlogComment
}

func NewCommonService() CommonService {
	return &commonService{
		dao: dao.NewCommonDao(datasource.InstanceMaster()),
	}
}

type commonService struct {
	dao *dao.CommonDao
}

func (s commonService) ListCommonByArticleId(articleId string) []models.BlogComment {
	return s.dao.ListCommonByArticleId(articleId, 10)
}

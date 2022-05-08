package service

import (
	"parent-api-go/internal/model/mop"
	"parent-api-go/pkg/app"
)

type ContentArticleListRequest struct {
	SubjectID int32 `form:"subject_id" binding:"required,min=0,max=100"`
	Status    uint8 `form:"status,default=1" binding:"oneof=0 1"`
}

func (svc *Service) GetContentArticleList(param *ContentArticleListRequest, pager *app.Pager, lineId uint8, fields []string, order string) (ArticleList []*mop.ContentArticle, ArticleCount int, err error) {
	ArticleCount, err = svc.daoMopSlave.CountArticleListBySubjectID(param.SubjectID, lineId)
	if err != nil {
		return
	}
	ArticleList, err = svc.daoMopSlave.GetArticleListBySubjectID(param.SubjectID, pager.Page, pager.PageSize, lineId, fields, order)
	return
}

func (svc *Service) GetArticleDetail(id int32) (article *mop.ContentArticle, err error) {
	article, err = svc.daoMopSlave.GetArticleDetailById(id)
	return
}

package dao

import (
	"parent-api-go/internal/model/mop"
)

func (d *Dao) CountArticleListBySubjectID(subjectID int32, lineId uint8) (int, error) {
	contentArticle := mop.ContentArticle{SubjectID: subjectID, LineID: lineId}
	return contentArticle.CountArticleListBySubjectID(d.engine)
}

func (d *Dao) GetArticleListBySubjectID(subjectID int32, offset, limit int, lineId uint8, fields []string, order string) ([]*mop.ContentArticle, error) {
	contentArticle := mop.ContentArticle{SubjectID: subjectID, LineID: lineId}
	return contentArticle.GetArticleListBySubjectID(d.engine, offset, limit, fields, order)
}

func (d *Dao) GetArticleDetailById(id int32) (*mop.ContentArticle, error) {
	contentArticle := mop.ContentArticle{}
	contentArticle.ID = id
	return contentArticle.GetArticleDetailById(d.engine)
}

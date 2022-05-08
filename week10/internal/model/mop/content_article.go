package mop

import (
	"github.com/jinzhu/gorm"
	"parent-api-go/internal/model"
	"time"
)

const StatusIsOnline = 1

var ArticleDefaultFields = []string{"id", "subject_id", "subject_name", "title", "subtitle", "cover", "content", "status", "created_at"}

type ContentArticle struct {
	*model.Model
	SubjectID   int32     `json:"-"`
	SubjectName string    `json:"-"`
	Title       string    `json:"title"`
	SubTitle    string    `json:"subtitle"`
	Cover       string    `json:"cover"`
	Content     string    `json:"-"`
	Status      uint8     `json:"status"`
	CreatedBy   int       `json:"-"`
	CreatedAt   time.Time `json:"-"`
	LineID      uint8     `json:"-"`
}

func (a *ContentArticle) TableName() string {
	return "app_content_article"
}

func (a *ContentArticle) CountArticleListBySubjectID(db *gorm.DB) (count int, err error) {
	err = db.Model(&a).Where("subject_id = ? AND status = ? ", a.SubjectID, StatusIsOnline).Count(&count).Error
	return
}
func (a *ContentArticle) GetArticleListBySubjectID(db *gorm.DB, offset int, limit int, fields []string, order string) (contentArticles []*ContentArticle, err error) {
	if len(fields) == 0 {
		fields = ArticleDefaultFields
	}

	db = db.Select(fields).Where("subject_id = ? AND status = ? ", a.SubjectID, StatusIsOnline)
	if a.LineID > 0 {
		db = db.Where("line_id = ?", a.LineID)
	}

	if order != "" {
		db = db.Order(order)
	}
	err = db.Offset(offset).Limit(limit).Find(&contentArticles).Error
	return
}

func (a *ContentArticle) GetArticleDetailById(db *gorm.DB) (*ContentArticle, error) {
	contentArticle := ContentArticle{}
	fields := []string{
		"id", "subject_name", "title", "subtitle", "cover", "created_at",
	}
	if err := db.Select(fields).Where("id = ?", a.ID).Find(&contentArticle).Error; err != nil {
		return nil, err
	}
	return &contentArticle, nil
}

package mop

import (
	"github.com/jinzhu/gorm"
	"time"
)

var SubjectDefaultFields = []string{"id", "name", "type", "status", "created_at"}
var SubjectTypeConfig = map[int]string{
	0: "",
	1: "图文",
	2: "视频",
}

type ContentSubject struct {
	ID        int32     `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	TypeShow  string    `json:"type_show"`
	Status    uint8     `json:"status"`
	CreatedBy int       `json:"created_by"`
	UpdatedBy int       `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
	Sort      int       `json:"sort" gorm:"default:1"`
}

func (*ContentSubject) TableName() string {
	return "app_content_subject"
}

func (cs *ContentSubject) GetSubjectById(db *gorm.DB, fields []string) (*ContentSubject, error) {
	if len(fields) == 0 {
		fields = SubjectDefaultFields
	}
	contentSubject := ContentSubject{}
	if err := db.Select(fields).Where("id = ?", cs.ID).Find(&contentSubject).Error; err != nil {
		return nil, err
	}
	return &contentSubject, nil
}

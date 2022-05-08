package mop

import (
	"github.com/jinzhu/gorm"
	"parent-api-go/internal/model"
	"time"
)

var VideoDefaultFields = []string{"id", "subject_id", "subject_name", "title",
	"subtitle", "video_cover", "video_url", "video_length_time", "student_name",
	"status", "created_at", "video_name", "line_id", "student_age", "study_length_time",
}

type ContentVideo struct {
	*model.Model
	SubjectID       int32     `json:"subject_id"`
	SubjectName     string    `json:"subject_name"`
	Title           string    `json:"title"`
	SubTitle        string    `json:"subtitle"`
	VideoCover      string    `json:"video_cover"`
	VideoUrl        string    `json:"video_url"`
	VideoLengthTime string    `json:"video_length_time"`
	Status          uint8     `json:"status"`
	StudentName     string    `json:"student_name"`
	StudentAge      string    `json:"student_age"`
	StudyLengthTime string    `json:"study_length_time"`
	CreatedBy       int       `json:"created_by"`
	CreatedAt       time.Time `json:"created_at"`
	VideoName       string    `json:"video_name"`
	LineID          uint8     `json:"line_id"`
}

func (a *ContentVideo) TableName() string {
	return "app_content_video"
}

func (a *ContentVideo) CountVideoListBySubjectID(db *gorm.DB) (count int, err error) {
	db = db.Model(&a).Where("subject_id = ? AND status = ? ", a.SubjectID, 1)
	if a.LineID > 0 {
		db = db.Where("line_id = ?", a.LineID)
	}
	err = db.Count(&count).Error
	return
}
func (a *ContentVideo) GetVideoListBySubjectID(db *gorm.DB, offset int, limit int, order string) (contentVideos []*ContentVideo, err error) {
	fields := VideoDefaultFields
	db = db.Select(fields).
		Where("subject_id = ? AND status = ? ", a.SubjectID, 1)

	if a.LineID > 0 {
		db = db.Where("line_id = ?", a.LineID)
	}

	if order != "" {
		db = db.Order(order)
	}

	err = db.Offset(offset).Limit(limit).Find(&contentVideos).Error
	return
}

func (a *ContentVideo) GetVideoDetailById(db *gorm.DB) (contentVideo *ContentVideo, err error) {
	fields := VideoDefaultFields
	err = db.Select(fields).Where("id = ?", a.ID).Find(&contentVideo).Error
	return
}

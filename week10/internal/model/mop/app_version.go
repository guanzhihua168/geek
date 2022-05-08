package mop

import "github.com/jinzhu/gorm"

const IsAuditing = 1
const IsAudited = 2

type AppVersionModel struct {
	ID         int    `json:"id"`
	Production int    `json:"production"`
	Platform   int    `json:"platform"`
	AppVersion string `json:"app_version"`
	IsAudition int    `json:"is_audition"`
}

func (a AppVersionModel) TableName() string {
	return "app_version"
}

func (a AppVersionModel) GetLatestVersionConfig(db *gorm.DB) (AppVersionModel, error) {
	var appVersion []AppVersionModel
	fields := []string{"id", "production", "platform", "app_version", "is_audition"}
	db = db.Select(fields)
	db = db.Where("platform = ? AND app_version = ?", a.Platform, a.AppVersion)
	db = db.Order("`app_version`.`id` DESC").Limit(1)
	err := db.Find(&appVersion).Error
	if err != nil || len(appVersion) == 0 {
		return AppVersionModel{}, err
	}

	return appVersion[0], nil
}

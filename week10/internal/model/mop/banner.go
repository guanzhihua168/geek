package mop

import (
	"github.com/jinzhu/gorm"
	"parent-api-go/internal/model"
	"parent-api-go/pkg/util"
)

const BannerMaxNumber = 5
const BannerIsOnline = 1
const BannerPosition = 2
const BannerForApp = 1
const AppV2Version = 2

type BannerModel struct {
	*model.Model
	Position       int    `json:"position"`
	Channel        string `json:"channel"`
	BannerName     string `json:"banner_name"`
	BannerImgUrl   string `json:"banner_img_url"`
	Link           string `json:"link"`
	Comment        string `json:"comment"`
	Status         uint8  `json:"status"`
	TargetUserType uint8  `json:"target_user_type"`
}

func (b BannerModel) TableName() string {
	return "banner"
}

func (b BannerModel) GetBannersToApp(db *gorm.DB, fields []string, version int, lineId uint8, limit int) (banners []*BannerModel, err error) {
	if limit == 0 {
		limit = BannerMaxNumber
	}

	nowDate := util.NowDate()
	err = db.Select(fields).
		Where("version = ? AND status = ? AND "+
			"position = ? AND platform = ? AND verify_status != ? AND line_id = ?",
			version, BannerIsOnline, BannerPosition, BannerForApp, IsAuditing, lineId).
		Where("(show_start_time < ? or show_start_time is null)", nowDate).
		Where("(show_end_time >= ? or show_end_time is null)", nowDate).
		Order("weight DESC").
		Order("id DESC").
		Limit(limit).
		Find(&banners).Error
	return
}

func (b BannerModel) GetBannersToAppAuditing(db *gorm.DB, fields []string, version int, lineId uint8, limit int) (banners []*BannerModel, err error) {
	if limit == 0 {
		limit = BannerMaxNumber
	}

	nowDate := util.NowDate()
	err = db.Select(fields).
		Where("version = ? AND status = ? AND "+
			"position = ? AND platform = ? AND verify_status = ? AND line_id = ?",
			version, BannerIsOnline, BannerPosition, BannerForApp, IsAuditing, lineId).
		Where("(show_start_time < ? or show_start_time is null)", nowDate).
		Where("(show_end_time >= ? or show_end_time is null)", nowDate).
		Order("weight DESC").
		Order("id DESC").
		Limit(limit).
		Find(&banners).Error
	return
}

func (b BannerModel) GetBannersToAppAudited(db *gorm.DB, fields []string, version int, lineId uint8, limit int) (banners []*BannerModel, err error) {
	if limit == 0 {
		limit = BannerMaxNumber
	}

	nowDate := util.NowDate()
	err = db.Select(fields).
		Where("version = ? AND status = ? AND "+
			"position = ? AND platform = ? AND verify_status = ? AND line_id = ?",
			version, BannerIsOnline, BannerPosition, BannerForApp, IsAuditing, lineId).
		Where("(show_start_time < ? or show_start_time is null)", nowDate).
		Where("(show_end_time >= ? or show_end_time is null)", nowDate).
		Order("weight DESC").
		Order("id DESC").
		Limit(limit).
		Find(&banners).Error
	return
}

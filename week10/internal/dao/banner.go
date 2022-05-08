package dao

import "parent-api-go/internal/model/mop"

func (d *Dao) GetBannersToApp(feilds []string, version int, lineId uint8, limit int) ([]*mop.BannerModel, error) {
	bannerModel := mop.BannerModel{}
	return bannerModel.GetBannersToApp(d.engine, feilds, version, lineId, limit)
}

func (d *Dao) GetBannersToAppAuditing(feilds []string, version int, lineId uint8, limit int) ([]*mop.BannerModel, error) {
	bannerModel := mop.BannerModel{}
	return bannerModel.GetBannersToAppAuditing(d.engine, feilds, version, lineId, limit)
}

func (d *Dao) GetBannersToAppAudited(feilds []string, version int, lineId uint8, limit int) ([]*mop.BannerModel, error) {
	bannerModel := mop.BannerModel{}
	return bannerModel.GetBannersToAppAudited(d.engine, feilds, version, lineId, limit)
}

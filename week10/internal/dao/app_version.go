package dao

import (
	"parent-api-go/internal/model/mop"
)

func (d *Dao) GetLatestVersionConfig(platform int, version string) (mop.AppVersionModel, error) {
	appVersion := mop.AppVersionModel{Platform: platform, AppVersion: version}
	return appVersion.GetLatestVersionConfig(d.engine)
}

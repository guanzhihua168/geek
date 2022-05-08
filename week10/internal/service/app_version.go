package service

import (
	"parent-api-go/internal/model/mop"
)

type AppVersionRequest struct {
	Platform   int    `form:"platform" binding:"required"`
	AppVersion string `form:"app_version" binding:"required"`
}

type AppVersion struct {
	ID         int    `json:"id"`
	Production int    `json:"production"`
	Platform   int    `json:"platform"`
	AppVersion string `json:"app_version"`
	IsAudition int    `json:"is_audition"`
}

func (svc *Service) GetLatestVersionConfig(platform int, version string) (mop.AppVersionModel, error) {
	appVersionConfig, err := svc.daoMopSlave.GetLatestVersionConfig(platform, version)
	if err != nil {
		return mop.AppVersionModel{}, err
	}
	return appVersionConfig, nil
}

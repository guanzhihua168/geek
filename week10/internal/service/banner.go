package service

import (
	"parent-api-go/global"
	"parent-api-go/internal/model/mop"
)

func (svc *Service) BannerFormat(data []*mop.BannerModel, userId uint32, limit int) (ret []*mop.BannerModel, err error) {
	if limit == 0 {
		limit = mop.BannerMaxNumber
	}

	var (
		userApmCount   int
		isUserLabelApm bool
	)

	for _, v := range data {
		if v.TargetUserType == global.IsImportUser || v.TargetUserType == global.IsLabelUser {
			if userId <= 0 {
				continue
			}

			// 人群包数据校验
			userApmCount, err = svc.daoMopSlave.IsUserApm(userId, v.ID, global.ApmTypeBanner)
			if v.TargetUserType == global.IsImportUser && userApmCount <= 0 {
				continue
			}

			// 用户标签校验
			isUserLabelApm, err = svc.IsUserLabelApm(userId, v.ID, nil)
			if v.TargetUserType == global.IsImportUser && !isUserLabelApm {
				continue
			}
		}

		v.BannerImgUrl = svc.GetImageUrl(v.BannerImgUrl, "")
		ret = append(ret, v)

		// 满足最大数量，跳出循环
		if len(ret) >= limit {
			break
		}
	}
	return
}

package service

import (
	"parent-api-go/global"
	"parent-api-go/internal/model/mop"
	"parent-api-go/internal/repository"
	"parent-api-go/pkg/util"
	"strings"
	"time"
)

/**
 * @Description: 获取需要显示的广告
 * @receiver svc
 * @param studentId 学生id
 * @param userId 用户id
 * @param isScreenAd 是否是屏幕广告
 * @param lineId 业务线
 * @param mobileOS 手机platform
 * @return *mop.PushMessageModel
 * @return error
 */
func (svc *Service) GetAdFirst(studentId, userId uint32, isScreenAd bool, lineId uint8, mobileOS uint8) (*mop.PushMessageModel, error) {
	dateTime := time.Now()
	var (
		err error
		ads []*mop.PushMessageModel
	)
	if isScreenAd {
		ads, err = svc.daoMopSlave.GetScreenAdByNowTime(dateTime, lineId)
	} else {
		ads, err = svc.daoMopSlave.GetAdByNowTime(dateTime, lineId)
	}

	if len(ads) == 0 || err != nil {
		return nil, err
	}

	// 闪屏广告用户未登录 & 投放用户是全体用户: 3
	if studentId <= 0 {
		if ads[0].TargetUserType == 3 {
			return ads[0], nil
		} else {
			return nil, nil
		}
	}

	// 选择返回某个广告
	return svc.selectAd(ads, studentId, userId, mobileOS)
}

/**
 * @Description: 根据不同目标人群选择广告
 * @receiver svc
 * @param ads
 * @param studentId
 * @param userId
 * @param mobileOS
 * @return *mop.PushMessageModel
 * @return error
 */
func (svc *Service) selectAd(ads []*mop.PushMessageModel, studentId, userId uint32, mobileOS uint8) (*mop.PushMessageModel, error) {
	var (
		err        error
		i          int
		r          bool
		isVip      bool
		userLabels []*repository.UserLabels
	)

	// 提前获取数据，防止循环多次获取
	firstAdTargetUT := ads[0].TargetUserType
	if firstAdTargetUT > 0 && firstAdTargetUT == global.IsImportUser || firstAdTargetUT == global.IsLabelUser {
		ulRepos := repository.UserLabelRepos{}
		ulRepos.Init()
		userLabels, err = ulRepos.GetUserLabel(userId)
	} else {
		// 是否购买过主课 todo 需要改为调用班级服务
		isVip = true
	}

	for _, v := range ads {
		// 平台不符合
		if v.Platform != global.PushOsAll && v.Platform != mobileOS {
			continue
		}

		if v.TargetUserType == global.PushTargetUserAll {
			// 全部用户
			return v, nil
		} else if v.TargetUserType == global.IsImportUser {
			// 人群包
			if i, err = svc.daoMopSlave.IsUserApm(userId, v.ID, global.ApmTypePush); i > 0 {
				return v, err
			}
		} else if v.TargetUserType == global.IsLabelUser {
			// 特定标签用户
			if r, err = svc.IsUserLabelApm(userId, v.ID, userLabels); r {
				return v, err
			}
		} else {
			if isVip {
				if v.TargetUserType == global.PushTargetUserIsVip {
					return v, nil
				}
			} else {
				if v.TargetUserType == global.PushTargetUserIsNotVip {
					return v, nil
				}
			}
		}
	}
	return nil, nil
}

/**
 * @Description: 校验用户标签是否已配置
 * @receiver svc
 * @param userId
 * @param messageId
 * @param userLabels
 * @return bool
 * @return error
 */
func (svc *Service) IsUserLabelApm(userId uint32, messageId int32, userLabels []*repository.UserLabels) (bool, error) {
	var (
		err          error
		relation     int8
		apmLabels    []*mop.ApmLabelModel
		userLabelIds []int
		apmLabelIds  []int
	)

	if userId <= 0 || messageId <= 0 {
		return false, nil
	}

	// 获取用户标签
	if len(userLabels) == 0 {
		ulRepos := repository.UserLabelRepos{}
		ulRepos.Init()
		userLabels, err = ulRepos.GetUserLabel(userId)
		if err != nil {
			return false, err
		}
	}

	// 获取配置标签
	apmLabels, err = svc.daoMopSlave.GetApmLabels(messageId, 0)
	if len(apmLabels) == 0 || err != nil {
		return false, err
	}

	// 标签关系
	if apmLabels[0].Relation > 0 {
		relation = 1
	} else {
		relation = 0
	}

	// 获取标签id
	userLabelIds = getUserLabelIds(userLabels)
	apmLabelIds = getApmLabelIds(apmLabels)

	// 取交集
	intersect := util.New(userLabelIds...).Intersect(util.New(apmLabelIds...))

	flag1 := len(userLabelIds) == len(apmLabelIds) && intersect.Count() == len(userLabelIds)
	flag2 := intersect.Count() > 0
	if relation > 0 {
		return flag1, nil
	} else {
		return flag2, nil
	}
}

/**
 * @Description: 格式化闪屏广告
 * @receiver *Service
 * @param ad
 * @return map[string]interface{}
 */
func (*Service) FormatScreenAd(ad *mop.PushMessageModel) map[string]interface{} {
	if ad == nil {
		return map[string]interface{}{}
	}

	requiredLogin := ad.TargetUserType != 3
	isClosed := ad.IsClosed > 0

	res := map[string]interface{}{
		"id":             ad.ID,
		"image":          ad.ImageUrl,
		"image_type":     "",
		"link":           ad.Link,
		"start_time":     ad.PushStartTime,
		"end_time":       ad.PushEndTime,
		"times":          ad.AdShowNumber,
		"sec":            ad.AdShowSec,
		"close":          isClosed,
		"required_login": requiredLogin,
		"line_id":        ad.LineId,
	}

	if ad.ImageUrl != "" {
		s := strings.Split(ad.ImageUrl, ".")
		res["image_type"] = s[len(s)-1]
	}

	return res
}

/**
 * @Description: 格式化弹屏广告
 * @receiver *Service
 * @param ad
 * @return map[string]interface{}
 */
func (*Service) FormatAd(ad *mop.PushMessageModel) map[string]interface{} {
	if ad == nil {
		return map[string]interface{}{}
	}

	if ad.LineId <= 0 {
		ad.LineId = 1
	}

	res := map[string]interface{}{
		"id":              ad.ID,
		"title":           ad.Title,
		"contents":        ad.Contents,
		"link":            ad.Link,
		"image_url":       ad.ImageUrl,
		"push_start_time": util.UnixToDate(ad.PushStartTime),
		"push_end_time":   util.UnixToDate(ad.PushEndTime),
		"ad_show_type":    ad.AdShowType,
		"ad_show_number":  ad.AdShowNumber,
		"line_id":         ad.LineId,
	}
	return res
}

/**
 * @Description: 获取用户标签ID
 * @param userLabels
 * @return []int
 */
func getUserLabelIds(userLabels []*repository.UserLabels) []int {
	var userLabelIds []int
	for _, v := range userLabels {
		userLabelIds = append(userLabelIds, int(v.TagId))
	}
	return userLabelIds
}

/**
 * @Description: 获取apm labelIds
 * @param apmLabels
 * @return []int
 */
func getApmLabelIds(apmLabels []*mop.ApmLabelModel) []int {
	var apmLabelIds []int
	for _, v := range apmLabels {
		apmLabelIds = append(apmLabelIds, int(v.LabelId))
	}
	return apmLabelIds
}

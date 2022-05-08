package service

import (
	"parent-api-go/pkg/cdn"
	"strings"
)

const CdnTypeParent = "parent"

/**
 * @Description: 获取图片访问地址
 * @param key
 * @param bucketType 项目bucket类型
 * @return string
 */
func (svc *Service) GetImageUrl(key, bucketType string) string {
	if key == "" {
		return ""
	}

	if bucketType == "" {
		bucketType = CdnTypeParent
	}

	// 兼容腾讯云地址
	if strings.HasPrefix(key, "http") {
		if strings.Contains(key, "myqcloud.com") {
			return cdn.GetQCloudCdnUrl(key, bucketType)
		}
		return key
	}
	return cdn.GetQiNiuCdnUrl(key)
}

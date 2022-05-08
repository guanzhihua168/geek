package cdn

import (
	"context"
	"net/url"
	"parent-api-go/global"
	"parent-api-go/pkg/setting"
	"strings"
)

/**
 * @Description: 获取资源cdn解析后地址 (myqcloud)
 * @param httpUrl 资源url地址
 * @param platformType 平台类型 如: tms, 配置文件 rouchi_conf/php_base/cdn.php
 * @return string
 */
func GetQCloudCdnUrl(httpUrl, platformType string) string {
	if httpUrl == "" || platformType == "" {
		return httpUrl
	}

	conf, ok := global.Cdn[platformType]
	if !ok {
		return httpUrl
	}

	urls, err := url.Parse(httpUrl)
	if err != nil {
		global.Logger.Warnf(context.Background(), "[GetQCloudCdnUrl]url.Parse(%s)fail!,err:%v", httpUrl, err)
		return httpUrl
	}

	// 屏蔽有些地址是相对路径
	host := urls.Host
	if host == "" {
		return httpUrl
	}

	_, ok2 := conf[host]
	if !ok2 {
		for _, v := range global.Cdn {
			if domain, ok3 := v[host]; ok3 && domain != "" {
				conf = v
				break
			}
		}
	}

	if domain2, ok4 := conf[host]; ok4 && domain2 != "" {
		// 过滤掉https或http
		httpUrl = strings.ReplaceAll(httpUrl, "https://", "")
		httpUrl = strings.ReplaceAll(httpUrl, "http://", "")

		// 替换域名
		httpUrl = strings.ReplaceAll(httpUrl, host, strings.Trim(domain2, "/"))
	}

	return httpUrl
}

/**
 * @Description: 获取七牛资源cdn解析后地址
 * @param key
 * @return string
 */
func GetQiNiuCdnUrl(key string) string {
	var domain string
	s, err := setting.NewQiNiuSetting()
	if err != nil {
		global.Logger.Errorf(context.Background(), "[GetQiNiuCdnUrl]setting.NewQiNiuSetting, err:%v", err)
		return ""
	}
	err = s.ReadSection("bucket.image.domain", &domain)
	if err != nil {
		global.Logger.Errorf(context.Background(), "[GetQiNiuCdnUrl]s.ReadSection, err:%v", err)
		return ""
	}
	return strings.TrimRight(domain, "/") + "/" + key
}

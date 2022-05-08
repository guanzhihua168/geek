package global

// 真实域名 => cdn域名
var Cdn = map[string]map[string]string{
	"tms":      {"jyxb-tms-1252525514.cos.ap-shanghai.myqcloud.com": "https://media-t.jingyuxiaoban.com"},
	"mop":      {"jyxb-common-1252525514.cos.ap-shanghai.myqcloud.com": "https://image-m.jingyuxiaoban.com"},
	"parent":   {"jyxb-common-1252525514.cos.ap-shanghai.myqcloud.com": "https://image-m.jingyupeiyou.com"},
	"playback": {"jyxb-playback-1252525514.cos.ap-shanghai.myqcloud.com": "http://media-playback.jingyuxiaoban.com"},
	"yunying":  {"jyxb-yunying-1252525514.cos.ap-shanghai.myqcloud.com": "https://yuying-t.jingyupeiyou.com"},

	// 上课平台
	"weclass_edu":     {"jyxb-weclass-edu-1252525514.cos.ap-shanghai.myqcloud.com": "https://media-e.jingyupeiyou.com"},
	"weclass_release": {"jyxb-weclass-release-1252525514.cos.ap-shanghai.myqcloud.com": "https://weclass-release-cdn.rouchi.com"},

	// 外包
	"wb_jyhb":        {"wb-jyhb-1252525514.cos.ap-shanghai.myqcloud.com": "https://wb-c.jingyupeiyou.com"},
	"parent_picture": {"jypy-zl-1252525514.cos.ap-shanghai.myqcloud.com": "https://jypy-zl.jingyupeiyou.com"},

	// tp端课件
	"tp_courseware": {"jypy-tp-courseware.oss-cn-shanghai-internal.aliyuncs.com": "https://jypy-tp-courseware.whalesenglish.com"},

	// 评测
	"pingce": {"jypy-pingce-1252525514.cos.ap-shanghai.myqcloud.com": "https://pingce-m.jingyupeiyou.com"},

	// m3u8 视频播放地址CDN
	"m3u8_video": {"jyxb-playback.oss-cn-shanghai.aliyuncs.com": "https://jypy-playback.jingyupeiyou.com"},

	// 发票服务相关
	"invoice": {"jypy-fapiao-1252525514.cos.ap-shanghai.myqcloud.com": "https://jypy-fapiao.jingyupeiyou.com"},
	"default": {"default": "https://image-m.jingyupeiyou.com"},
}

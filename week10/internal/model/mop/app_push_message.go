package mop

import (
	"github.com/jinzhu/gorm"
	"parent-api-go/global"
	"parent-api-go/internal/model"
	"parent-api-go/pkg/util"
	"time"
)

/*
  `sub_title` varchar(255) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT 'push 消息副标题',
  `image_url` varchar(500) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '图片原地址',
  `link` varchar(500) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '跳转链接',
  `platform` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '平台 {1: ios, 2:andorid;3:all}',
  `target_user_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '推送用户类型 {1:vip; 2:unvip;3:all,4在读,5非再读 6购买过主课 7未购买主课}',
  `send_numbers` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '准备要推送给的用户数',
  `reach_numbers` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '显示此刻送达的人数，已结束的推送显示推送已结束后送达的人数',
  `open_numbers` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '打开人数',
  `push_start_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '指定推送开始时间',
  `send_status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '发送状态{1:未预览;2:已预览,待发送;3:预览完成,可发送;4:发送中;5:成功发送;6:取消发送}',
  `createtime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updatetime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `operator_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '操作人uid',
  `operator_username` varchar(50) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '操作者姓名',
  `expire_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '过期时间',
  `contents` varchar(500) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '站内消息内容',
  `message_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '站内消息类型{1: 系统; 2:活动}',
  `style_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '样式类型{0: 纯文本,1:大图, 2: 小图}',
  `is_join_android_task` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否加入android推送任务列表 0. 未加入 1.已加入',
  `is_join_ios_task` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否加入iso推送任务',
  `push_course_id` varchar(400) DEFAULT '0' COMMENT '发送的课程id 未0则未选择课程',
  `push_course_type` varchar(255) DEFAULT NULL COMMENT '选择课程类型 多个使用,',
  `image_type` tinyint(20) unsigned DEFAULT '0' COMMENT '图片类型 【0 无图  1 大图 2 大图】',
  `course_is_class` tinyint(3) unsigned DEFAULT '0' COMMENT '当天是否有课(0 无课 1 有课)',
  `course_class_start_time` int(11) DEFAULT NULL COMMENT '当前有课推送的开始时间',
  `course_class_expire_time` int(11) DEFAULT NULL COMMENT '当天有课推送的结束时间',
  `push_end_time` int(11) unsigned DEFAULT NULL COMMENT 'push_type=3 显示的到期时间',
  `is_broad_online` tinyint(255) unsigned DEFAULT '0',
  `is_closed` tinyint(3) unsigned DEFAULT '0' COMMENT '是否可关闭(push_type=3的时候才存在)',
  `ad_show_type` tinyint(3) unsigned DEFAULT '1' COMMENT '(push_type=4) 每人每天最多展示几次',
  `ad_show_number` tinyint(3) unsigned DEFAULT '0' COMMENT '(push_type=4)每人每天展示几次',
  `ad_show_sec` int(11) DEFAULT '0' COMMENT '广告展示秒数 push_type=5',
*/
type PushMessageModel struct {
	*model.Model
	Title          string `json:"title"`
	SubTitle       string `json:"sub_title"`
	PushType       uint8  `json:"push_type"`
	AdStatus       uint8  `json:"ad_status"`
	PushStartTime  int64  `json:"push_start_time"`
	PushEndTime    int64  `json:"push_end_time"`
	LineId         uint8  `json:"line_id"`
	TargetUserType uint8  `json:"target_user_type"`
	Platform       uint8  `json:"platform"`
	ImageUrl       string `json:"image_url"`
	Link           string `json:"link"`
	AdShowNumber   uint8  `json:"ad_show_number"`
	AdShowType     uint8  `json:"ad_show_type"`
	AdShowSec      int    `json:"ad_show_sec"`
	IsClosed       uint8  `json:"is_closed"`
	Contents       string `json:"contents"`
}

func (*PushMessageModel) TableName() string {
	return "app_push_message"
}

func (pm *PushMessageModel) GetScreenAdByNowTime(db *gorm.DB, time time.Time) (pushMessages []*PushMessageModel, err error) {
	dateTime := util.FormatDate(time)
	db = db.Where("push_type = ? AND ad_status = ? "+
		"AND push_start_time <= ? AND push_end_time >= ?",
		global.AdScreenMessageType, global.AdSendStatus, dateTime, dateTime)
	if pm.LineId > 0 {
		db = db.Where("line_id = ?", pm.LineId)
	}
	err = db.Order("id DESC").Find(&pushMessages).Error
	return
}

func (pm *PushMessageModel) GetAdByNowTime(db *gorm.DB, time time.Time) (pushMessages []*PushMessageModel, err error) {
	dateTime := util.FormatDate(time)
	db = db.Where("push_type = ? AND ad_status = ? "+
		"AND push_start_time <= ? AND push_end_time >= ?",
		global.AdMessageType, global.AdSendStatus, dateTime, dateTime)
	if pm.LineId > 0 {
		db = db.Where("line_id = ?", pm.LineId)
	}
	err = db.Order("id DESC").Find(&pushMessages).Error
	return
}

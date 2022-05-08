package mop

import (
	"github.com/jinzhu/gorm"
	"parent-api-go/global"
	"parent-api-go/internal/model"
	"time"
)

type ApmUserModel struct {
	*model.Model
	Type      uint8     `json:"type"`
	MessageId int32     `json:"message_id"`
	UserId    uint32    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (*ApmUserModel) TableName() string {
	return "apm_user"
}

// 用户是否配置
func (am *ApmUserModel) IsUserApm(db *gorm.DB) (count int, err error) {
	err = db.Where("message_id", am.MessageId).
		Where("user_id", am.UserId).
		Where("type", global.ActiveEnable).Count(&count).Error
	return
}

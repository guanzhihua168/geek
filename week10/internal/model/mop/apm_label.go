package mop

import (
	"github.com/jinzhu/gorm"
	"parent-api-go/global"
	"parent-api-go/internal/model"
	"time"
)

type ApmLabelModel struct {
	*model.Model
	Type      uint8     `json:"type"`
	MessageId int32     `json:"message_id"`
	LabelId   int32     `json:"label_id"`
	Relation  uint8     `json:"relation"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (*ApmLabelModel) TableName() string {
	return "apm_label"
}

func (am *ApmLabelModel) GetApmLabels(db *gorm.DB) (apmLabels []*ApmLabelModel, err error) {
	err = db.Where("message_id = ? AND type = ? AND active = ?",
		am.MessageId, am.Type, global.ActiveEnable).Find(&apmLabels).Error
	return
}

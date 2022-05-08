package dao

import (
	"parent-api-go/global"
	"parent-api-go/internal/model/mop"
)

func (d *Dao) GetApmLabels(messageId int32, Type uint8) ([]*mop.ApmLabelModel, error) {
	if Type == 0 {
		Type = global.ApmTypePush
	}
	apmLabel := mop.ApmLabelModel{MessageId: messageId, Type: Type}
	return apmLabel.GetApmLabels(d.engine)
}

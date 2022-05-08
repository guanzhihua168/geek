package dao

import "parent-api-go/internal/model/mop"

func (d *Dao) IsUserApm(userId uint32, messageId int32, Type uint8) (int, error) {
	apmUserModel := mop.ApmUserModel{UserId: userId, MessageId: messageId, Type: Type}
	return apmUserModel.IsUserApm(d.engine)
}

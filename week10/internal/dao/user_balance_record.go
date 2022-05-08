package dao

import "parent-api-go/internal/model"

func (d *Dao) GetParentBalanceTransFlowRecords(userID uint32, recordType int) ([]*model.UserBalanceRecordModel, error) {
	balanceRecordModel := model.UserBalanceRecordModel{UserID: userID, Type: recordType}
	return balanceRecordModel.GetParentBalanceTransFlowRecords(d.engine)
}

func (d *Dao) GetParentBalanceTransFlowRecordDetail(recordID int) (*model.UserBalanceRecordModel, error) {
	balanceRecordModel := model.UserBalanceRecordModel{}
	return balanceRecordModel.GetParentBalanceTransFlowRecordDetail(d.engine)
}

package model

import "github.com/jinzhu/gorm"

type UserBalanceRecordModel struct {
	*Model
	UserID    uint32 `json:"user_id"`
	BalanceID int    `json:"balance_id"`
	Balance   int    `json:"balance"`
	Fee       int    `json:"fee"`
	Status    int    `json:"status"`
	Type      int    `json:"type"`
	TradeType int    `json:"trade_type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (m UserBalanceRecordModel) TableName() string {
	return "user_balance_record"
}

//获取某个用户的帐户余额
func (m UserBalanceRecordModel) GetParentBalanceTransFlowRecords(db *gorm.DB) ([]*UserBalanceRecordModel, error) {
	var records []*UserBalanceRecordModel
	fields := []string{
		"id", "user_id", "balance_id", "balance", "type", "fee", "trade_type", "status", "active", "created_at",
	}
	db = db.Table(m.TableName()).Select(fields).Where("user_id= ? and type = ? active = ?", m.UserID, m.Type, 1)
	if err := db.Scan(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (m UserBalanceRecordModel) GetParentBalanceTransFlowRecordDetail(db *gorm.DB) (*UserBalanceRecordModel, error) {
	var record *UserBalanceRecordModel

	fields := []string{
		"id", "user_id", "balance_id", "balance", "type", "fee", "trade_type", "status", "active", "created_at",
	}
	if err := db.Select(fields).Where("id = ? and active = ?", m.ID, 1).Find(&record).Error; err != nil {
		return nil, err
	}

	return record, nil

}

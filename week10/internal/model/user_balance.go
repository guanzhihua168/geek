package model

import (
	"github.com/jinzhu/gorm"
)

type UserBalanceModel struct {
	*Model
	UserID    uint32 `json:"user_id"`
	Balance   int    `json:"balance"`
	Status    int    `json:"status"`
	Type      int    `json:"type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (b UserBalanceModel) TableName() string {
	return "user_balance"
}

//获取某个用户的帐户余额
func (b UserBalanceModel) GetTotalBalanceByUserID(db *gorm.DB) (int, error) {
	var userBalance UserBalanceModel
	db = db.Table(b.TableName()).Select("sum(balance) as balance").Where("user_id = ?", b.UserID)
	if err := db.Scan(&userBalance).Error; err != nil {
		return 0, err
	}

	return userBalance.Balance, nil
}

//func (b UserBalanceModel) getUserBalanceByType(db *gorm.DB) ([]*UserBalanceModel, error)  {
//	var userBalance []*UserBalanceModel
//}

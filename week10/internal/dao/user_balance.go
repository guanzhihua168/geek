package dao

import "parent-api-go/internal/model"

func (d *Dao) GetTotalMoneyBalanceByUserID(userID uint32) (int, error) {
	UserBalanceModel := model.UserBalanceModel{UserID: userID}
	return UserBalanceModel.GetTotalBalanceByUserID(d.engine)
}

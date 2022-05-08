package dao

import (
	"parent-api-go/internal/model/mop"
	"time"
)

func (d *Dao) GetScreenAdByNowTime(time time.Time, lineId uint8) ([]*mop.PushMessageModel, error) {
	pushModel := mop.PushMessageModel{LineId: lineId}
	return pushModel.GetScreenAdByNowTime(d.engine, time)
}

func (d *Dao) GetAdByNowTime(time time.Time, lineId uint8) ([]*mop.PushMessageModel, error) {
	pushModel := mop.PushMessageModel{LineId: lineId}
	return pushModel.GetAdByNowTime(d.engine, time)
}

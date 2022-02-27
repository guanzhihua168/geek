package pkg

import (
	"geek/week5/global"
	"time"
)

type limiting struct {
	cnt        uint
	timeWindow int64
}

// 单机滑动窗口限流
func (l *limiting) LimitReq(limitKey string) bool {
	currT := time.Now().Unix()

	if global.LimitQ == nil {
		global.LimitQ = make(map[string][]int64, 0)
	}

	if _, ok := global.LimitQ[limitKey]; !ok {
		global.LimitQ[limitKey] = append(global.LimitQ[limitKey], currT)
	}

	if uint(len(global.LimitQ[limitKey])) < l.cnt {
		return true
	}

	earliestT := global.LimitQ[limitKey][0]
	if currT-earliestT <= l.timeWindow {
		return false
	}

	global.LimitQ[limitKey] = global.LimitQ[limitKey][1:]
	global.LimitQ[limitKey] = append(global.LimitQ[limitKey], currT)

	return true
}

func NewLimiting() *limiting {
	return &limiting{cnt: global.DefaultCnt, timeWindow: global.DefaultTw}
}

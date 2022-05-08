package repository

import (
	"encoding/json"
	"parent-api-go/global"
	"parent-api-go/internal/model"
	"parent-api-go/pkg/linkerd"
	"parent-api-go/pkg/util"
	"time"

	"github.com/sirupsen/logrus"
)

type OrderRepos struct {
	repository
}

func (c *OrderRepos) Init() {
	options := linkerdOptions(c.repository)
	options = append(options, linkerd.WithJson())

	c.Remote = linkerd.NewLinkerd(
		global.LinkerdSetting.Host,
		global.LinkerdSetting.AppName,
		global.LinkerdSetting.Token,
		"rouchi_order",
		options...,
	)
}

// 根据学生ID、开始时间、结束时间、类型 获取订单
// https://yapi.rouchi.com/project/294/interface/api/12753
func (c *OrderRepos) GetOrderPaySuccess(uid uint32, startTime, endTime time.Time, orderType, status []int, active int) (order []model.Order, err error) {

	data := struct {
		Uid          []uint32 `json:"uid"`
		PayTimeStart string   `json:"pay_time_start"`
		PayTimeEnd   string   `json:"pay_time_end"`
		Active       []int    `json:"active"`
		OrderType    []int    `json:"order_type"`
		Status       []int    `json:"status"`
	}{
		[]uint32{uid},
		util.FormatDate(startTime),
		util.FormatDate(endTime),
		[]int{active},
		orderType,
		status,
	}

	b, _ := json.Marshal(data)
	j, e := c.Remote.Post("/v2/order/getByQuery", b)

	if e != nil {
		logrus.Warning(e.Error())
		return []model.Order{}, e
	}

	if e = json.Unmarshal(j.Data, &order); e != nil || len(order) == 0 {
		return []model.Order{}, e
	}

	return
}

package service

import (
	"parent-api-go/global"
	"parent-api-go/internal/model"
	"parent-api-go/internal/repository"
	"parent-api-go/pkg/util"
	"time"
)

func (svc *Service) BatchGetUserPaidOrderBySourcePrice(userId uint32, orderType int, sources []int) (order []model.Order, err error) {
	var startPaidTime time.Time
	// 订单状态
	orderStatus := []int{global.OrderPaymentDone}
	orderRepos := repository.OrderRepos{}
	orderRepos.Init()
	startPaidTime, err = util.TimeLocalFmt(global.InvoiceOrderStartPaidTime)
	if err != nil {
		return order, err
	}
	order, err = orderRepos.GetOrderPaySuccess(userId, startPaidTime, time.Now(), []int{orderType}, orderStatus, global.ActiveEnable)
	if err != nil {
		return order, err
	}

	result := order[:0]
	for _, o := range order {
		if util.InArrayInts(o.SourceId, sources) {
			result = append(result, o)
		}
	}
	return result, nil
}

func (svc *Service) GetUserPaidOrderFilterPrice(order []model.Order, price int32) []model.Order {
	result := order[:0]
	for _, o := range order {
		if o.RevenuePrice >= price {
			result = append(result, o)
		}
	}
	return result
}

// 二维slice排序
func (svc *Service) SliceMultiSortForOrderModel(orders []model.Order) (newOrders []model.Order, err error) {
	var (
		paidAt1, paidAt2 time.Time
		ordersLen        int
	)

	ordersLen = len(orders)
	for i := 0; i < ordersLen; i++ {
		for j := 0; j < ordersLen-i-1; j++ {
			orderJ := orders[j]
			paidAt1, err = util.TimeLocalFmt(orderJ.PaidAt)
			if err != nil {
				return newOrders, err
			}

			orderJ2 := orders[j+1]
			paidAt2, err = util.TimeLocalFmt(orderJ2.PaidAt)
			if err != nil {
				return newOrders, err
			}

			if paidAt1.Unix() < paidAt2.Unix() {
				orders[j], orders[j+1] = orders[j+1], orders[j]
			}

		}
	}

	return orders, err
}

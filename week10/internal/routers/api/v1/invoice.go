package v1

import (
	"parent-api-go/global"
	"parent-api-go/internal/model"
	"parent-api-go/internal/service"
	"parent-api-go/pkg/app"
	"parent-api-go/pkg/context"
	"parent-api-go/pkg/errcode"
	"parent-api-go/pkg/util"
	"strconv"
)

type Invoice struct {
}

func NewInvoice() Invoice {
	return Invoice{}
}

func (Invoice) GetUserInvoiceList(c *context.AppContext) {

	response := app.NewResponse(c.Context)
	orderType, err := strconv.Atoi(c.Context.Query("order_type"))
	if err != nil {
		response.ToErrorResponse(errcode.ServerError)
		return
	}

	// 允许的订单类型。1课程订单、3充值订单
	if util.InArrayInts(orderType, []int{global.OrderTypeCourse, global.OrderTypeTopUp}) {
		response.ToErrorResponse(errcode.InvalidOrderType)
		return
	}

	svc := service.New(c.Request.Context())
	orders, e := svc.GetUserInvoiceOrderList(c.AuthId, orderType)
	if e != nil {
		global.Logger.Errorf(c.Context, "[invoice][getUserInvoiceOrderList]error,userId:%d", c.AuthId)
		response.ToResponse(nil)
	}

	if len(orders) == 0 {
		response.ToResponse(nil)
	}

	// 返回值处理
	response.ToResponse(formatInvoiceOrderList(orders))
}

type InvoiceResult struct {
	Id       uint32 `json:"id"`
	No       string `json:"no"`
	Amount   string `json:"amount"`
	PaidTime string `json:"paid_time"`
}

func formatInvoiceOrderList(orders []model.Order) (results []InvoiceResult) {
	for _, v := range orders {
		tmpV := InvoiceResult{
			Id:       v.ID,
			No:       v.No,
			Amount:   util.NumberFormat(strconv.Itoa(int(v.RevenuePrice) / 100)),
			PaidTime: v.PaidAt,
		}
		results = append(results, tmpV)
	}
	return results
}

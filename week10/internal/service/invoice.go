package service

import (
	"github.com/tidwall/gjson"
	"parent-api-go/global"
	"parent-api-go/internal/model"
	"parent-api-go/internal/repository"
	"parent-api-go/pkg/util"
	"strconv"
)

/**
* @Description: 获取用户可开票订单列表
   *     订单来源属于：官网微信端、官网WAP端、官网PC端、iOS客户端、安卓客户端、BOSS、CRM，共7个来源的，无论金额多少，均展示；并可开具发票。
   *     订单来源属于：天猫平台、京东平台、淘宝平台、有赞平台、外部定制系统、小鹅通平台、小哈皮平台、其他平台，共8个来源的，金额大于等于40元的订单；并可开具发票。
* @receiver svc
* @param userId 用户id
* @param orderType 订单类型
*/
func (svc *Service) GetUserInvoiceOrderList(userId uint32, orderType int) (order []model.Order, err error) {
	if userId <= 0 || orderType <= 0 {
		return
	}

	order, err = svc.BatchGetUserPaidOrderBySourcePrice(userId, orderType, global.InvoiceNoPriceSources)
	if err != nil {
		return
	}
	// 以下来源不需要考虑付款金额
	unpaidOrder := svc.GetUserPaidOrderFilterPrice(order, global.OrderSourceNoPrice)
	// 以下来源金额需不低于40元
	paidOrder := svc.GetUserPaidOrderFilterPrice(order, global.OrderSourcePrice)

	// 合并结果集
	ret := mergeOrder(unpaidOrder, paidOrder)

	// 获取可开票的订单id
	orderNos := getOrderNos(ret)
	validOrderNos, e := svc.BatchGetValidInvoiceOrderNo(userId, orderNos)
	if e != nil {
		return order, e
	}

	var newRet []model.Order
	// 过滤不可开票的订单
	for _, v := range ret {
		if util.InArrayStrings(v.No, validOrderNos) {
			newRet = append(newRet, v)
		}
	}

	// 按照付款时间倒序排列
	return svc.SliceMultiSortForOrderModel(newRet)
}

func (svc *Service) BatchGetValidInvoiceOrderNo(userId uint32, orderNos []string) (validOrderNos []string, err error) {
	if userId <= 0 || len(orderNos) == 0 {
		return
	}
	var (
		invoices    []repository.InvoiceDetail
		newOrderNos []string
	)
	invoices, err = svc.BatchGetOrderInvoiceList(userId, orderNos)
	if len(invoices) == 0 {
		return orderNos, err
	}

	// 筛选出已申请、已开票的订单编号
	for _, invoice := range invoices {
		if invoice.InvoiceStatus == global.StatusProcess || invoice.InvoiceStatus == global.StatusFinish {
			for _, v := range invoice.OrderNo {
				if !util.InArrayStrings(v, newOrderNos) {
					newOrderNos = append(newOrderNos, v)
				}
			}
		}
	}

	// 过滤已申请、已开票订单编号
	for _, orderNo := range orderNos {
		if !util.InArrayStrings(orderNo, newOrderNos) {
			validOrderNos = append(validOrderNos, orderNo)
		}
	}
	return
}

/**
 * @Description: 批量获取订单开票信息
 * @param userId 用户id
 * @param orderNos 订单编号
 */
func (svc *Service) BatchGetOrderInvoiceList(userId uint32, orderNos []string) (newRet []repository.InvoiceDetail, err error) {
	if userId <= 0 || len(orderNos) == 0 {
		return
	}

	// 分组获取: 允许订单最大个数为50
	bigOrderNos := util.SliceChunkForString(orderNos, 50)
	var (
		ret        []repository.InvoiceDetail
		ok         bool
		jsonRaw    string
		invoiceRaw repository.InvoiceWithRaw
	)

	kMap := make(map[int]map[string]interface{})

	invoiceRepos := repository.InvoiceRepos{}
	invoiceRepos.Init()
	for k0, chunkOrderNos := range bigOrderNos {
		invoiceRaw, err = invoiceRepos.GetInvoicesByOrderNos(chunkOrderNos)
		if err != nil {
			break
		}
		if len(invoiceRaw.InvoiceDetail) > 0 {
			ret = mergeInvoiceDetail(ret, invoiceRaw.InvoiceDetail)
			kMap[k0] = map[string]interface{}{
				"ret":     ret,
				"jsonRaw": invoiceRaw.RetRaw,
			}
		}
	}

	if len(ret) == 0 {
		return
	}

	// 处理返回值pdf地址、发票抬头
	for _, v := range kMap {
		ret, ok = v["ret"].([]repository.InvoiceDetail)
		if !ok {
			continue
		}
		jsonRaw, ok = v["jsonRaw"].(string)
		if !ok {
			continue
		}

		for k, vv := range ret {
			formatInvoiceDetail(jsonRaw, k, &vv)
			newRet = append(newRet, vv)
		}
	}
	return

}

/**
 * @Description: 发票服务-发票详情相关接口返回值格式化
 * @param invoiceDetail
 * @return repository.InvoiceDetail
 */
func formatInvoiceDetail(retRaw string, k int, invoiceDetail *repository.InvoiceDetail) {
	getJson := func(kStr string) gjson.Result {
		return gjson.Get(retRaw, strconv.Itoa(k)+"."+kStr)
	}

	// 开票状态
	invoiceStatusR := getJson("invoice_status")
	statusR := getJson("status")
	if invoiceStatusR.Exists() {
		if invoiceDetail.InvoiceStatus <= 0 {
			invoiceDetail.InvoiceStatus = global.StatusProcess
		}
	} else if statusR.Exists() {
		if invoiceDetail.Status <= 0 {
			invoiceDetail.Status = global.StatusProcess
		}
	}

	//发票抬头类型
	invoiceHeaderTypeR := getJson("invoice_header_type")
	headerTypeR := getJson("header_type")
	if invoiceHeaderTypeR.Exists() {
		if invoiceDetail.InvoiceHeaderType == 1 {
			invoiceDetail.HeaderType = global.TitleTypeCorp
		} else {
			invoiceDetail.HeaderType = global.TitleTypePersonal
		}
	} else if headerTypeR.Exists() {
		if invoiceDetail.HeaderType == 1 {
			invoiceDetail.HeaderType = global.TitleTypeCorp
		} else {
			invoiceDetail.HeaderType = global.TitleTypePersonal
		}
	}

	// 发票pdf链接
	invoiceDetail.InvoicePdf = parsePdfUrl(invoiceDetail.InvoicePdf)
}

func parsePdfUrl(invoicePdf string) string {
	return gjson.Get(invoicePdf, "url").String()
}

func getOrderNos(order []model.Order) (nos []string) {
	for _, v := range order {
		nos = append(nos, v.No)
	}
	return nos
}

func mergeOrder(orders ...[]model.Order) (order []model.Order) {
	for _, v := range orders {
		for _, vv := range v {
			order = append(order, vv)
		}
	}
	return order
}

func mergeInvoiceDetail(invoiceDetails ...[]repository.InvoiceDetail) (invoiceDetail []repository.InvoiceDetail) {
	for _, v := range invoiceDetails {
		for _, vv := range v {
			invoiceDetail = append(invoiceDetail, vv)
		}
	}
	return invoiceDetail
}

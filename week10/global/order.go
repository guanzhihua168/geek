package global

// 支付成功
const OrderPaymentDone = 2

//订单类型
const OrderTypeCourse = 1
const OrderTypeTopUp = 3

//订单来源
const OrderSourceOfficialWechat = 1
const OrderSourceOfficialWap = 2
const OrderSourceOfficialPc = 3
const OrderSourceIos = 4
const OrderSourceAndroid = 5
const OrderSourceBoss = 6
const OrderSourceCrm = 7
const OrderSourceTMall = 101
const OrderSourceJD = 102
const OrderSourceTaoBao = 103
const OrderSourceYouZan = 104
const OrderSourceOutsource = 105
const OrderSourceXiaoETong = 106
const OrderSourceHapi = 108
const OrderSourceOther = 999

//发票校验使用，指定来源金额不低于该金额
const OrderSourceNoPrice = 100
const OrderSourcePrice = 4000

//发票允许开票订单开始时间
const InvoiceOrderStartPaidTime = "2020-07-09 00:00:00"

//发票校验使用，订单来源组合：以下来源不需要考虑金额
var InvoiceNoPriceSources = []int{
	OrderSourceOfficialWechat,
	OrderSourceOfficialWap,
	OrderSourceOfficialPc,
	OrderSourceIos,
	OrderSourceAndroid,
	OrderSourceBoss,
	OrderSourceCrm,
}

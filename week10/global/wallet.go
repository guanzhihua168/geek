package global

const PcpuTypeGeneralClass int = 5 //通用课时包
const PcpuTypeAllPower int = 12    //全能课时包

//商品cpu每节课时长
var CPUDuration = map[int]string{
	0:  "0分钟",
	25: "25分钟",
	50: "50分钟",
	60: "60分钟",
}

const PcpuPurchase int = 1
const PcpuBookClassroom int = 2
const PcpuDropClassroom int = 3
const PcpuRetire int = 4
const PcpuGift int = 5
const PcpuTransfer int = 6
const PcpuBookPyp int = 7
const PcpuCanceledPyp int = 8
const PcpuDismissClassroom int = 9
const PcpuSwitchTextBookFrom int = 10
const PcpuSwitchTextbookTo int = 11
const PcpuRecordTypeRefundConsumption int = 20 //退课消
const PcpuRecordTypeFriendSell int = 21        //转介绍赠送
const PcpuRecordTypeManualGift int = 22        //手动赠送
const PcpuRecordTypeManualCancel int = 23      //手动扣减
const PcpuRecordTypeSwitchTypeFrom int = 24    //转出通用课时包
const PcpuRecordTypeSwitchTypeTo int = 25      //转入通用课时包
const PcpuRecordTypeAutoCancel int = 27        //自动退课扣减
const PcpuRecordTypePointsCharge int = 28      //积分兑换

var PcpuBalancePlusTypeList = []int{
	PcpuPurchase,
	PcpuDropClassroom,
	PcpuGift,
	PcpuTransfer, // 转换通用课时，文案改为退出班级，便于家长理解
	PcpuCanceledPyp,
	PcpuDismissClassroom,
	PcpuSwitchTextbookTo,
	PcpuRecordTypeRefundConsumption,
	PcpuRecordTypeFriendSell,
	PcpuRecordTypeManualGift,
	PcpuRecordTypeManualCancel,
	PcpuRecordTypeSwitchTypeFrom,
	PcpuRecordTypeSwitchTypeTo,
	PcpuRecordTypeAutoCancel,
	PcpuRecordTypePointsCharge,
}

const PcpuBalanceIoPlus int = 1
const PcpuBalanceIoMinus int = 2

const UserBalanceTradeTypeSystemAdd = 0
const UserBalanceTradeTypeRecharge = 1
const UserBalanceTradeTypeCancel = 2
const UserBalanceTradeTypeExitClassroom = 3
const UserBalanceTradeTypeGift = 4

//用户账户交易类型
const UserBalanceTradeTypePcpuRefund = 6
const UserBalanceTradeTypeRecommend = 7
const UserBalanceTradeTypeCollectLikes = 8
const UserBalanceTradeTypeMakeUp = 9 //补偿金
const UserBalanceTradeTypeEvent = 10 //活动赠送
const UserBalanceTradeTypeSystemDeduct = 100
const UserBalanceTradeTypePay = 101
const UserBalanceTradeTypeDeduct = 102
const UserBalanceTradeTypeWithdraw = 103

//用户账户交易类型。5奖励
const UserBalanceTradeTypeBonus = 5

var UserBalanceTradeType = map[int]string{
	UserBalanceTradeTypeSystemAdd:     "系统增加",
	UserBalanceTradeTypeRecharge:      "缴费充值",
	UserBalanceTradeTypeCancel:        "取消订单",
	UserBalanceTradeTypeExitClassroom: "取消班级订单",
	UserBalanceTradeTypeGift:          "赠送",
	UserBalanceTradeTypeBonus:         "奖励",
	UserBalanceTradeTypePcpuRefund:    "退课时包",
	UserBalanceTradeTypeRecommend:     "推荐有礼",
	UserBalanceTradeTypeCollectLikes:  "集赞有礼",
	UserBalanceTradeTypeMakeUp:        "补偿金",
	UserBalanceTradeTypeEvent:         "活动赠送",

	UserBalanceTradeTypeSystemDeduct: "系统减少",
	UserBalanceTradeTypePay:          "支付",
	UserBalanceTradeTypeDeduct:       "抵扣",
	UserBalanceTradeTypeWithdraw:     "退费",
}

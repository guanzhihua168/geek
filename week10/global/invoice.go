package global

//发票类型。1电子发票 2纸质发票
const TypeEInvoice = 1
const TypePaperInvoice = 2

//发票抬头默认。1是 0否
const TitleDefault = 1
const TitleNotDefault = 0

//发票抬头类型。1个人 2企业
const TitleTypePersonal = 1
const TitleTypeCorp = 2

//开票状态。5开票中，10已开票，15开票失败
const StatusProcess = 5
const StatusFinish = 10
const StatusFail = 15

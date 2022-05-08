package main

type a struct {
}

func NewA() a {
	return a{}
}

func main() {
	//t,_ := util.TimeLocalFmt("2021-01-01 08:09:00")
	//fmt.Println(t.Format("2006-01-02 15:04:05"))
	//s := []int{1,2,3}
	//s2 := []int{8,2,3}
	//r := util.New(s...).Intersect(util.New(s2...))
	//fmt.Println(r.List())
	//tm := time.Unix(1574677687, 0)
	//fmt.Println(util.FormatDate(tm))
	//u, _ := url.Parse("https://media-e.jingyupeiyou.com/adf/a?a=1")

	//fmt.Println(u.Host)
	//s := cdn.GetQCloudCdnUrl("https://jyxb-tms-1252525514.cos.ap-shanghai.myqcloud.com/adf/a?b=2", "weclass_edu")
	//startPaidTime, err := util.TimeLocalFmt(global.InvoiceOrderStartPaidTime)
	//s := []string{"a","b","c","d","e","f","g","h","i","j","k"}
	//ss := util.SliceChunkForString(s, 4)
	//fmt.Println(ss)
	//j := "{\"order_no\":\"232324242\", \"invoice_status\":0}"
	//invoice := repository.InvoiceDetail{}
	//err := json.Unmarshal([]byte(j), &invoice)
	//
	//fmt.Println(invoice.InvoiceStatus, err)

	//orders := []model.Order{
	//	{PaidAt: "2021-06-07 11:40:00"},
	//	{PaidAt: "2021-06-07 11:39:00"},
	//	{PaidAt: "2021-06-07 12:39:00"},
	//	{PaidAt: "2021-06-07 10:39:00"},
	//	{PaidAt: "2021-06-07 12:09:00"},
	//}
	//
	//svc := service.New(context.Background())
	//fmt.Println(svc.SliceMultiSortForOrderModel(orders))

}

type Ss struct {
	K string
	V int
}

func formatS(v *Ss) {
	if v.V == 2 {
		v.V = 4
	}
}

package repository

import (
	"encoding/json"
	"net/url"
	"parent-api-go/global"
	"parent-api-go/pkg/linkerd"
)

type InvoiceRepos struct {
	repository
}

func (c *InvoiceRepos) Init() {
	options := linkerdOptions(c.repository)
	options = append(options, linkerd.WithJson())

	c.Remote = linkerd.NewLinkerd(
		global.LinkerdSetting.Host,
		global.LinkerdSetting.AppName,
		global.LinkerdSetting.Token,
		"service_invoice",
		options...,
	)
}

type InvoiceDetail struct {
	OrderNo           []string `json:"order_no"`
	InvoiceStatus     int      `json:"invoice_status,omitempty"`
	Status            int      `json:"status,omitempty"`
	InvoiceHeaderType int      `json:"invoice_header_type,omitempty"`
	HeaderType        int      `json:"header_type"`
	InvoicePdf        string   `json:"invoice_pdf"`
	InvoiceFailReason string   `json:"invoice_fail_reason"`
}

type InvoiceWithRaw struct {
	InvoiceDetail []InvoiceDetail
	RetRaw        string `json:"-"`
}

func (c *InvoiceRepos) GetInvoicesByOrderNos(orderNos []string) (invoiceRaw InvoiceWithRaw, err error) {
	data := &url.Values{}
	for _, v := range orderNos {
		data.Add("orderNos[]", v)
	}

	j, e := c.Remote.Post("/V2/Invoice/getInvoicesByOrderNos", []byte(data.Encode()))
	if e != nil {
		c.Ctx.Log.Warning(e.Error())
		return invoiceRaw, nil
	}

	invoiceRaw.RetRaw = string(j.Data)

	if err = json.Unmarshal(j.Data, &invoiceRaw.InvoiceDetail); err != nil || len(invoiceRaw.InvoiceDetail) == 0 {
		invoiceRaw.InvoiceDetail = []InvoiceDetail{}
		return
	}

	return
}

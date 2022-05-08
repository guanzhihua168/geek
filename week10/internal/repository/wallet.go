package repository

import (
	"encoding/json"
	"fmt"
	"parent-api-go/global"
	"parent-api-go/pkg/linkerd"
	"time"
)

type WalletRepos struct {
	repository
}
type WalletBalance struct {
	Total                   float64 `json:"total"`
	ChargeBalanceTotal      float64 `json:"chargeBalanceTotal"`
	ScholarshipBalanceTotal float64 `json:"scholarshipBalanceTotal"`
}
type ResultData struct {
	Pcpus   []Pcpus      `json:"pcpus"`
	Records []PcpuRecord `json:"records"`
}
type Pcpus struct {
	ID                         int         `json:"id"`
	No                         string      `json:"no"`
	Type                       int         `json:"type"`
	Status                     int         `json:"status"`
	Capacity                   int         `json:"capacity"`
	Active                     int         `json:"active"`
	Duration                   int         `json:"duration"`
	TypeFmt                    string      `json:"type_fmt"`
	StatusFmt                  string      `json:"status_fmt"`
	StudentID                  int         `json:"student_id"`
	OrderID                    interface{} `json:"order_id"`
	ParentPcpuID               interface{} `json:"parent_pcpu_id"`
	RestSessionTimes           int         `json:"rest_session_times"`
	CPUID                      interface{} `json:"cpu_id"`
	OriginalPrice              int         `json:"original_price"`
	ActualPrice                int         `json:"actual_price"`
	OriginalSingleSessionPrice int         `json:"original_single_session_price"`
	ActualSingleSessionPrice   int         `json:"actual_single_session_price"`
	SessionTimes               int         `json:"session_times"`
	CourseID                   interface{} `json:"course_id"`
	CourseType                 int         `json:"course_type"`
	CourseLevelFrom            int         `json:"course_level_from"`
	CourseLvelTo               int         `json:"course_lvel_to"`
	JdProductID                interface{} `json:"jd_product_id"`
	TmallProductID             interface{} `json:"tmall_product_id"`
	CreatedAt                  string      `json:"created_at"`
	UpdatedAt                  string      `json:"updated_at"`
	RealPrice                  int         `json:"real_price"`
	RealSingleSessionPrice     int         `json:"real_single_session_price"`
	YouzanProductID            interface{} `json:"youzan_product_id"`
	MiniRealPrice              int         `json:"mini_real_price"`
	MiniRealSingleSessionPrice int         `json:"mini_real_Single_session_price"`
	MinimumPrice               int         `json:"minimum_price"`
	MinimumSingleSessionPrice  int         `json:"minimum_single_session_price"`
	NameCn                     string      `json:"name_cn"`
	NameEn                     string      `json:"name_en"`
	PaymentPrice               int         `json:"payment_price"`
	PaymentSingleSessionPrice  int         `json:"payment_single_session_price"`
	RevenuePrice               int         `json:"revenue_price"`
	RevenueSingleSessionPrice  int         `json:"revenue_single_session_price"`
	SalePrice                  int         `json:"sale_price"`
	SaleSingleSessionPrice     int         `json:"sale_single_session_price"`
	CourseTextbookType         int         `json:"course_textbook_type"`
}
type PcpuRecord struct {
	ID               int         `json:"id"`
	Type             int         `json:"type"`
	Memo             string      `json:"memo"`
	PcpuID           int         `json:"pcpu_id"`
	RestSessionTimes int         `json:"rest_session_times"`
	SessionTimes     int         `json:"session_times"`
	ClassroomID      interface{} `json:"classroom_id"`
	CreatedAt        time.Time   `json:"created_at"`
	TypeFmt          string      `json:"type_fmt"`
}

func (s *WalletRepos) Init() {
	options := linkerdOptions(s.repository)
	options = append(options, linkerd.WithJson())
	s.Remote = linkerd.NewLinkerd(
		global.LinkerdSetting.Host,
		global.LinkerdSetting.AppName,
		global.LinkerdSetting.Token,
		"rouchi_wallet",
		options...,
	)
}

func (s *WalletRepos) GetParentBalance(userId uint32) WalletBalance {
	s.Init()

	data := struct {
		UserID uint32 `json:"id"`
	}{
		userId,
	}

	body, _ := json.Marshal(data)
	j, e := s.Remote.Post("/v1/balance/parent/get", body)
	if e != nil {
		s.Ctx.Log.Warning(e.Error())
		return WalletBalance{}
	}

	b := WalletBalance{}
	_ = json.Unmarshal(j.Data, &b)
	return b
}

func (s *WalletRepos) GetStudentPcpus(studentID uint32) ([]Pcpus, error) {
	s.Init()

	data := struct {
		StudentID uint32 `json:"student_id"`
	}{
		studentID,
	}

	var resultData ResultData
	body, _ := json.Marshal(data)
	j, e := s.Remote.Post("/v1/pcpu/student/getWithRecord", body)
	if e != nil {
		s.Ctx.Log.Warning(fmt.Sprintf("%+v", e))
		return resultData.Pcpus, e
	}

	err := json.Unmarshal(j.Data, &resultData)
	if err != nil {
		return resultData.Pcpus, err
	}
	return resultData.Pcpus, nil
}

func (s *WalletRepos) GetStudentPcpuRecords(studentID uint32) ([]PcpuRecord, error) {
	s.Init()

	data := struct {
		StudentID  uint32 `json:"student_id"`
		WithRecord int    `json:"with_record"`
	}{
		studentID,
		1,
	}

	var resultData ResultData
	body, _ := json.Marshal(data)
	res, err := s.Remote.Post("/v1/pcpu/student/getWithRecord", body)
	if err != nil {
		s.Ctx.Log.Warning(fmt.Sprintf("%+v", err))
		return nil, err
	}

	err = json.Unmarshal(res.Data, &resultData)
	if err != nil {
		return nil, err
	}
	return resultData.Records, nil
}

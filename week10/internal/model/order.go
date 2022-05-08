package model

type Order struct {
	ID           uint32 `json:"id"`
	UserID       uint32 `json:"user_id"`
	StudentID    uint32 `json:"student_id"`
	No           string `json:"no"`
	PaidAt       string `json:"paid_at"`
	RevenuePrice int32  `json:"revenue_price"`
	Status       int    `json:"status"`
	SourceId     int    `json:"source_id"`

	OrderItems []OrderItem `json:"order_items"`
}

type OrderItem struct {
	ID      uint32 `json:"id"`
	OrderID uint32 `json:"order_id"`
	CpuID   uint32 `json:"cpu_id"`
}

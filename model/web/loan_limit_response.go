package web

type LoanLimitResponse struct {
	// Fields
	LoanLimitID  uint    `json:"loan_limit_id"`
	CustomerNik  string  `json:"customer_nik"`
	CustomerName string  `json:"customer_name"`
	OneMonth     float32 `json:"one_month"`
	TwoMonth     float32 `json:"two_month"`
	ThreeMonth   float32 `json:"three_month"`
	FourMonth    float32 `json:"four_month"`
}

package web

type LoanLimitUpdateRequest struct {
	// Fields
	OneMonth   float32 `json:"one_month"`
	TwoMonth   float32 `json:"two_month"`
	ThreeMonth float32 `json:"three_month"`
	FourMonth  float32 `json:"four_month"`
}

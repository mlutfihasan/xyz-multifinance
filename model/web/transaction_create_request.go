package web

type TransactionCreateRequest struct {
	// Fields
	CustomerNik    string  `json:"customer_nik"`
	CustomerName   string  `json:"customer_name"`
	OnTheRoad      float32 `json:"on_the_road"`
	AdminFee       float32 `json:"admin_fee"`
	LoanAmount     float32 `json:"loan_amount"`
	InterestAmount float32 `json:"interest_amount"`
	AssetName      string  `json:"asset_name"`
}

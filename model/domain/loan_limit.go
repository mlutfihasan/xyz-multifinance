package domain

import (
	"xyz-multifinance/model/web"
)

type LoanLimits []LoanLimit
type LoanLimit struct {
	// Fields
	LoanLimitID  uint    `gorm:"primarykey;autoIncrement;"`
	CustomerNik  string  `gorm:"size:16;not null;"`
	CustomerName string  `gorm:"size:60;not null;"`
	OneMonth     float32 `gorm:"not null;"`
	TwoMonth     float32 `gorm:"not null;"`
	ThreeMonth   float32 `gorm:"not null;"`
	FourMonth    float32 `gorm:"not null;"`
}

func (loanLimit *LoanLimit) ToLoanLimitResponse() web.LoanLimitResponse {
	return web.LoanLimitResponse{
		// Fields
		LoanLimitID:  loanLimit.LoanLimitID,
		CustomerNik:  loanLimit.CustomerNik,
		CustomerName: loanLimit.CustomerName,
		OneMonth:     loanLimit.OneMonth,
		TwoMonth:     loanLimit.TwoMonth,
		ThreeMonth:   loanLimit.ThreeMonth,
		FourMonth:    loanLimit.FourMonth,
	}
}

func (loanLimits LoanLimits) ToLoanLimitResponses() []web.LoanLimitResponse {
	loanLimitResponses := []web.LoanLimitResponse{}
	for _, loanLimit := range loanLimits {
		loanLimitResponses = append(loanLimitResponses, loanLimit.ToLoanLimitResponse())
	}
	return loanLimitResponses
}

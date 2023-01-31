package models

import (
	"time"

	"gorm.io/gorm"
)

type LoanLimitHistory struct {
	LoanId       int     `json:"loan_id"`
	CustomerNik  string  `json:"customer_nik"`
	CustomerName string  `json:"customer_name"`
	OneMonth     float32 `json:"one_month"`
	TwoMonth     float32 `json:"two_month"`
	ThreeMonth   float32 `json:"three_month"`
	FourMonth    float32 `json:"four_month"`

	ActionTaken string    `json:"action_taken"`
	ActionTime  time.Time `json:"action_time"`
}

func (llh *LoanLimitHistory) Prepare(action string, data LoanLimit) {
	llh.LoanId = data.LoanId
	llh.CustomerNik = data.CustomerNik
	llh.CustomerName = data.CustomerName
	llh.OneMonth = data.OneMonth
	llh.TwoMonth = data.TwoMonth
	llh.ThreeMonth = data.ThreeMonth
	llh.FourMonth = data.FourMonth

	llh.ActionTaken = action
	llh.ActionTime = time.Now()
}

func SaveLoanLimitHistory(db *gorm.DB, data LoanLimitHistory) error {
	err := db.Debug().Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type TransactionHistory struct {
	TransactionNo   string    `json:"transaction_no"`
	TransactionDate time.Time `json:"transaction_date"`
	CustomerNik     string    `json:"customer_nik"`
	CustomerName    string    `json:"customer_name"`
	OnTheRoad       float32   `json:"on_the_road"`
	AdminFee        float32   `json:"admin_fee"`
	LoanAmount      float32   `json:"loan_amount"`
	InterestAmount  float32   `json:"interest_amount"`
	AssetName       string    `json:"asset_name"`

	ActionTaken string    `json:"action_taken"`
	ActionTime  time.Time `json:"action_time"`
}

func (th *TransactionHistory) Prepare(action string, data Transaction) {
	th.TransactionNo = data.TransactionNo
	th.TransactionDate = data.TransactionDate
	th.CustomerNik = data.CustomerNik
	th.CustomerName = data.CustomerName
	th.OnTheRoad = data.OnTheRoad
	th.AdminFee = data.AdminFee
	th.LoanAmount = data.LoanAmount
	th.InterestAmount = data.InterestAmount
	th.AssetName = data.AssetName

	th.ActionTaken = action
	th.ActionTime = time.Now()
}

func SaveTransactionHistory(db *gorm.DB, data TransactionHistory) error {
	err := db.Debug().Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}

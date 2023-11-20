package domain

import (
	"time"
	"xyz-multifinance/model/web"
)

type Transactions []Transaction
type Transaction struct {
	// Fields
	TransactionNo   string    `gorm:"primarykey;size:50;not null;"`
	TransactionDate time.Time `gorm:"not null;"`
	CustomerNik     string    `gorm:"size:16;not null;"`
	CustomerName    string    `gorm:"size:60;not null;"`
	OnTheRoad       float32   `gorm:"not null;"`
	AdminFee        float32   `gorm:"not null;"`
	LoanAmount      float32   `gorm:"not null;"`
	InterestAmount  float32   `gorm:"not null;"`
	AssetName       string    `gorm:"size:50;not null;"`
}

func (transaction *Transaction) ToTransactionResponse() web.TransactionResponse {
	return web.TransactionResponse{
		// Fields
		TransactionNo:   transaction.TransactionNo,
		TransactionDate: transaction.TransactionDate,
		CustomerNik:     transaction.CustomerNik,
		CustomerName:    transaction.CustomerName,
		OnTheRoad:       transaction.OnTheRoad,
		AdminFee:        transaction.AdminFee,
		LoanAmount:      transaction.LoanAmount,
		InterestAmount:  transaction.InterestAmount,
		AssetName:       transaction.AssetName,
	}
}

func (transactions Transactions) ToTransactionResponses() []web.TransactionResponse {
	transactionResponses := []web.TransactionResponse{}
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, transaction.ToTransactionResponse())
	}
	return transactionResponses
}

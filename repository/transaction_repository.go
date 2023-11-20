package repository

import (
	"xyz-multifinance/model/domain"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindAll(db *gorm.DB, filters *map[string]string) domain.Transactions
	Create(db *gorm.DB, transaction *domain.Transaction) *domain.Transaction
	Update(db *gorm.DB, transactionNo *string, transaction *domain.Transaction) *domain.Transaction
	Delete(db *gorm.DB, transactionNo *string)
}

package repository

import (
	"xyz-multifinance/helper"
	"xyz-multifinance/model/domain"

	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
}

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (repository *TransactionRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.Transactions {
	transactions := domain.Transactions{}
	tx := db.Model(&domain.Transaction{})

	err := helper.ApplyFilter(tx, filters)
	helper.PanicIfError(err)

	err = tx.Find(&transactions).Error
	helper.PanicIfError(err)

	return transactions
}

func (repository *TransactionRepositoryImpl) Create(db *gorm.DB, transaction *domain.Transaction) *domain.Transaction {
	err := db.Create(&transaction).Error
	helper.PanicIfError(err)
	return transaction
}

func (repository *TransactionRepositoryImpl) Update(db *gorm.DB, transactionNo *string, transaction *domain.Transaction) *domain.Transaction {
	err := db.Model(&domain.Transaction{}).
		Where(&domain.Transaction{
			TransactionNo: *transactionNo,
		}).
		Updates(&transaction).Error
	helper.PanicIfError(err)

	err = db.First(&transaction, &domain.Transaction{TransactionNo: *transactionNo}).Error
	helper.PanicIfError(err)

	return transaction
}

func (repository *TransactionRepositoryImpl) Delete(db *gorm.DB, transactionNo *string) {
	transaction := &domain.Transaction{}
	tx := db.First(transaction, &domain.Transaction{TransactionNo: *transactionNo}).Updates(&domain.Transaction{
		TransactionNo: *transactionNo,
	})

	// Deleting the Transaction from the database.
	err := tx.Unscoped().Delete(transaction, &domain.Transaction{TransactionNo: *transactionNo}).Error
	helper.PanicIfError(err)
}

package service

import (
	"strconv"
	"time"
	"xyz-multifinance/helper"
	"xyz-multifinance/model/domain"
	"xyz-multifinance/model/web"
	"xyz-multifinance/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"gorm.io/gorm"
)

type TransactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
	DB                    *gorm.DB
	Validate              *validator.Validate
}

func NewTransactionService(
	transaction repository.TransactionRepository,
	db *gorm.DB,
	validate *validator.Validate,
) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepository: transaction,
		DB:                    db,
		Validate:              validate,
	}
}

func (service *TransactionServiceImpl) FindAll(filters *map[string]string) []web.TransactionResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	transactions := service.TransactionRepository.FindAll(tx, filters)
	return transactions.ToTransactionResponses()
}

func (service *TransactionServiceImpl) Create(request *web.TransactionCreateRequest, c *gin.Context) web.TransactionResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	err = tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	unixTimeNow := strconv.Itoa(int(time.Now().Unix()))
	transactionNo := request.CustomerNik + "/" + unixTimeNow

	transaction := &domain.Transaction{
		// Fields
		TransactionNo:   transactionNo,
		TransactionDate: time.Now(),
		CustomerNik:     request.CustomerNik,
		CustomerName:    request.CustomerName,
		OnTheRoad:       request.OnTheRoad,
		AdminFee:        request.AdminFee,
		LoanAmount:      request.LoanAmount,
		InterestAmount:  request.InterestAmount,
		AssetName:       request.AssetName,
	}

	transaction = service.TransactionRepository.Create(tx, transaction)

	return transaction.ToTransactionResponse()
}

func (service *TransactionServiceImpl) Update(transactionNo *string, request *web.TransactionUpdateRequest, c *gin.Context) web.TransactionResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	err = tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaction := &domain.Transaction{
		//  Fields
		OnTheRoad:      request.OnTheRoad,
		AdminFee:       request.AdminFee,
		LoanAmount:     request.LoanAmount,
		InterestAmount: request.InterestAmount,
		AssetName:      request.AssetName,
	}
	transaction = service.TransactionRepository.Update(tx, transactionNo, transaction)

	return transaction.ToTransactionResponse()
}

func (service *TransactionServiceImpl) Delete(transactionNo *string) {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	service.TransactionRepository.Delete(tx, transactionNo)
}

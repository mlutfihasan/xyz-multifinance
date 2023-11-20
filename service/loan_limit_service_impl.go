package service

import (
	"xyz-multifinance/helper"
	"xyz-multifinance/model/domain"
	"xyz-multifinance/model/web"
	"xyz-multifinance/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"gorm.io/gorm"
)

type LoanLimitServiceImpl struct {
	LoanLimitRepository repository.LoanLimitRepository
	DB                  *gorm.DB
	Validate            *validator.Validate
}

func NewLoanLimitService(
	loanLimit repository.LoanLimitRepository,
	db *gorm.DB,
	validate *validator.Validate,
) LoanLimitService {
	return &LoanLimitServiceImpl{
		LoanLimitRepository: loanLimit,
		DB:                  db,
		Validate:            validate,
	}
}

func (service *LoanLimitServiceImpl) FindAll(filters *map[string]string) []web.LoanLimitResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	loanLimits := service.LoanLimitRepository.FindAll(tx, filters)
	return loanLimits.ToLoanLimitResponses()
}

func (service *LoanLimitServiceImpl) Create(request *web.LoanLimitCreateRequest, c *gin.Context) web.LoanLimitResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	err = tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	loanLimit := &domain.LoanLimit{
		// Fields
		CustomerNik:  request.CustomerNik,
		CustomerName: request.CustomerName,
		OneMonth:     request.OneMonth,
		TwoMonth:     request.TwoMonth,
		ThreeMonth:   request.ThreeMonth,
		FourMonth:    request.FourMonth,
	}

	loanLimit = service.LoanLimitRepository.Create(tx, loanLimit)

	return loanLimit.ToLoanLimitResponse()
}

func (service *LoanLimitServiceImpl) Update(loanLimitID *string, request *web.LoanLimitUpdateRequest, c *gin.Context) web.LoanLimitResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	err = tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	loanLimit := &domain.LoanLimit{
		//  Fields
		OneMonth:   request.OneMonth,
		TwoMonth:   request.TwoMonth,
		ThreeMonth: request.ThreeMonth,
		FourMonth:  request.FourMonth,
	}
	loanLimit = service.LoanLimitRepository.Update(tx, loanLimitID, loanLimit)

	return loanLimit.ToLoanLimitResponse()
}

func (service *LoanLimitServiceImpl) Delete(loanLimitID *string) {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	service.LoanLimitRepository.Delete(tx, loanLimitID)
}

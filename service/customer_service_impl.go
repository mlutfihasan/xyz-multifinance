package service

import (
	"errors"
	"xyz-multifinance/helper"
	"xyz-multifinance/model/domain"
	"xyz-multifinance/model/web"
	"xyz-multifinance/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"gorm.io/gorm"
)

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	DB                 *gorm.DB
	Validate           *validator.Validate
}

func NewCustomerService(
	customer repository.CustomerRepository,
	db *gorm.DB,
	validate *validator.Validate,
) CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: customer,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *CustomerServiceImpl) FindAll(filters *map[string]string) []web.CustomerResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	customers := service.CustomerRepository.FindAll(tx, filters)
	return customers.ToCustomerResponses()
}

func (service *CustomerServiceImpl) Create(request *web.CustomerCreateRequest, c *gin.Context) web.CustomerResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	err = tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	var linkIdPhoto string
	idPhoto, _ := request.IdPhoto.Open()
	if idPhoto != nil && request.IdPhoto != nil {
		linkIdPhoto, err = helper.UploadFilePhoto(idPhoto, *request.IdPhoto, request.CustomerNik, helper.FileTypeIdPhoto)
		if err != nil {
			defer idPhoto.Close()
			helper.PanicIfError(err)
		}
		defer idPhoto.Close()
	} else {
		helper.PanicIfError(errors.New("required id photo"))
	}

	var linkSelfiePhoto string
	selfiePhoto, _ := request.SelfiePhoto.Open()
	if selfiePhoto != nil && request.SelfiePhoto != nil {
		linkSelfiePhoto, err = helper.UploadFilePhoto(selfiePhoto, *request.SelfiePhoto, request.CustomerNik, helper.FileTypeSelfiePhoto)
		if err != nil {
			defer selfiePhoto.Close()
			helper.PanicIfError(err)
		}
		defer selfiePhoto.Close()
	} else {
		helper.PanicIfError(errors.New("required id photo"))
	}

	customer := &domain.Customer{
		// Fields
		CustomerNik:       request.CustomerNik,
		CustomerName:      request.CustomerName,
		CustomerLegalName: request.CustomerLegalName,
		PlaceOfBirth:      request.PlaceOfBirth,
		DateOfBirth:       request.DateOfBirth,
		Salary:            request.Salary,
		IdPhoto:           linkIdPhoto,
		SelfiePhoto:       linkSelfiePhoto,
	}

	customer = service.CustomerRepository.Create(tx, customer)

	return customer.ToCustomerResponse()
}

func (service *CustomerServiceImpl) Update(customerNik *string, request *web.CustomerUpdateRequest, c *gin.Context) web.CustomerResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	err = tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer := &domain.Customer{
		//  Fields
		Salary: request.Salary,
	}
	customer = service.CustomerRepository.Update(tx, customerNik, customer)

	return customer.ToCustomerResponse()
}

func (service *CustomerServiceImpl) Delete(customerNik *string) {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	service.CustomerRepository.Delete(tx, customerNik)
}

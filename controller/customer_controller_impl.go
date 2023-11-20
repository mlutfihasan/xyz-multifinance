package controller

import (
	"net/http"
	"strconv"
	"time"

	"xyz-multifinance/helper"
	"xyz-multifinance/model/web"
	"xyz-multifinance/service"

	"github.com/gin-gonic/gin"
)

type CustomerControllerImpl struct {
	CustomerService service.CustomerService
}

func NewCustomerController(customerService service.CustomerService) CustomerController {
	return &CustomerControllerImpl{
		CustomerService: customerService,
	}
}

func (controller *CustomerControllerImpl) FindAll(c *gin.Context) {
	filters := helper.FilterFromQueryString(c)
	customerResponses := controller.CustomerService.FindAll(&filters)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(customerResponses),
		Data:    customerResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *CustomerControllerImpl) Create(c *gin.Context) {
	customerNik := c.PostForm("customer_nik")
	customerName := c.PostForm("customer_name")
	customerLegalName := c.PostForm("customer_legal_name")
	placeOfBirth := c.PostForm("place_of_birth")
	dateOfBirth := c.PostForm("date_of_birth")
	salary := c.PostForm("salary")

	idPhoto, err := c.FormFile("id_photo")
	if err != nil {
		helper.PanicIfError(err)
	}

	selfiePhoto, err := c.FormFile("selfie_photo")
	if err != nil {
		helper.PanicIfError(err)
	}

	layout := "2006-01-02T15:04:05Z"
	dateOfBirthTime, err := time.Parse(layout, dateOfBirth)
	if err != nil {
		helper.PanicIfError(err)
	}

	salaryFloat32, err := strconv.ParseFloat(salary, 32)
	if err != nil {
		helper.PanicIfError(err)
	}

	request := web.CustomerCreateRequest{
		CustomerNik:       customerNik,
		CustomerName:      customerName,
		CustomerLegalName: customerLegalName,
		PlaceOfBirth:      placeOfBirth,
		DateOfBirth:       dateOfBirthTime,
		Salary:            float32(salaryFloat32),
		IdPhoto:           idPhoto,
		SelfiePhoto:       selfiePhoto,
	}

	customerResponse := controller.CustomerService.Create(&request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Customer created successfully",
		Data:    customerResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *CustomerControllerImpl) Update(c *gin.Context) {
	customerNik := c.Param("customer_nik")

	request := web.CustomerUpdateRequest{}
	helper.ReadFromRequestBody(c, &request)

	customerResponse := controller.CustomerService.Update(&customerNik, &request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Customer created successfully",
		Data:    customerResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *CustomerControllerImpl) Delete(c *gin.Context) {
	customerNik := c.Param("customer_nik")

	controller.CustomerService.Delete(&customerNik)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Customer deleted successfully",
	}

	c.JSON(http.StatusOK, webResponse)
}

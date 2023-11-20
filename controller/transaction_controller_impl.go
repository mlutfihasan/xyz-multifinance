package controller

import (
	"net/http"

	"xyz-multifinance/helper"
	"xyz-multifinance/model/web"
	"xyz-multifinance/service"

	"github.com/gin-gonic/gin"
)

type TransactionControllerImpl struct {
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &TransactionControllerImpl{
		TransactionService: transactionService,
	}
}

func (controller *TransactionControllerImpl) FindAll(c *gin.Context) {
	filters := helper.FilterFromQueryString(c)
	transactionResponses := controller.TransactionService.FindAll(&filters)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(transactionResponses),
		Data:    transactionResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *TransactionControllerImpl) Create(c *gin.Context) {
	request := web.TransactionCreateRequest{}
	helper.ReadFromRequestBody(c, &request)

	transactionResponse := controller.TransactionService.Create(&request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Transaction created successfully",
		Data:    transactionResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *TransactionControllerImpl) Update(c *gin.Context) {
	transactionNo := c.Param("transaction_no")

	request := web.TransactionUpdateRequest{}
	helper.ReadFromRequestBody(c, &request)

	transactionResponse := controller.TransactionService.Update(&transactionNo, &request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Transaction created successfully",
		Data:    transactionResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *TransactionControllerImpl) Delete(c *gin.Context) {
	transactionNo := c.Param("transaction_no")

	controller.TransactionService.Delete(&transactionNo)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Transaction deleted successfully",
	}

	c.JSON(http.StatusOK, webResponse)
}

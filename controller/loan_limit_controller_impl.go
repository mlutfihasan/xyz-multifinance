package controller

import (
	"net/http"

	"xyz-multifinance/helper"
	"xyz-multifinance/model/web"
	"xyz-multifinance/service"

	"github.com/gin-gonic/gin"
)

type LoanLimitControllerImpl struct {
	LoanLimitService service.LoanLimitService
}

func NewLoanLimitController(loanLimitService service.LoanLimitService) LoanLimitController {
	return &LoanLimitControllerImpl{
		LoanLimitService: loanLimitService,
	}
}

func (controller *LoanLimitControllerImpl) FindAll(c *gin.Context) {
	filters := helper.FilterFromQueryString(c)
	loanLimitResponses := controller.LoanLimitService.FindAll(&filters)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(loanLimitResponses),
		Data:    loanLimitResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *LoanLimitControllerImpl) Create(c *gin.Context) {
	request := web.LoanLimitCreateRequest{}
	helper.ReadFromRequestBody(c, &request)

	loanLimitResponse := controller.LoanLimitService.Create(&request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Loan Limit created successfully",
		Data:    loanLimitResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *LoanLimitControllerImpl) Update(c *gin.Context) {
	loanLimitID := c.Param("loan_limit_id")

	request := web.LoanLimitUpdateRequest{}
	helper.ReadFromRequestBody(c, &request)

	loanLimitResponse := controller.LoanLimitService.Update(&loanLimitID, &request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Loan Limit created successfully",
		Data:    loanLimitResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *LoanLimitControllerImpl) Delete(c *gin.Context) {
	loanLimitID := c.Param("loan_limit_id")

	controller.LoanLimitService.Delete(&loanLimitID)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Loan Limit deleted successfully",
	}

	c.JSON(http.StatusOK, webResponse)
}

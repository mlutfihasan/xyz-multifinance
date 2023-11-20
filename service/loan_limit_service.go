package service

import (
	"xyz-multifinance/model/web"

	"github.com/gin-gonic/gin"
)

type LoanLimitService interface {
	FindAll(filters *map[string]string) []web.LoanLimitResponse
	Create(requests *web.LoanLimitCreateRequest, c *gin.Context) web.LoanLimitResponse
	Update(loanLimitID *string, requests *web.LoanLimitUpdateRequest, c *gin.Context) web.LoanLimitResponse
	Delete(loanLimitID *string)
}

package service

import (
	"xyz-multifinance/model/web"

	"github.com/gin-gonic/gin"
)

type TransactionService interface {
	FindAll(filters *map[string]string) []web.TransactionResponse
	Create(requests *web.TransactionCreateRequest, c *gin.Context) web.TransactionResponse
	Update(transactionNo *string, requests *web.TransactionUpdateRequest, c *gin.Context) web.TransactionResponse
	Delete(transactionNo *string)
}

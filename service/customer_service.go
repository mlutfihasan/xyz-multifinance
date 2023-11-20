package service

import (
	"xyz-multifinance/model/web"

	"github.com/gin-gonic/gin"
)

type CustomerService interface {
	FindAll(filters *map[string]string) []web.CustomerResponse
	Create(requests *web.CustomerCreateRequest, c *gin.Context) web.CustomerResponse
	Update(customerNik *string, requests *web.CustomerUpdateRequest, c *gin.Context) web.CustomerResponse
	Delete(customerNik *string)
}

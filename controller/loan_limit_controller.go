package controller

import (
	"github.com/gin-gonic/gin"
)

type LoanLimitController interface {
	FindAll(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

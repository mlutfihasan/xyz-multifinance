package app

import (
	"xyz-multifinance/helper"
	"xyz-multifinance/route"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, validate *validator.Validate) *gin.Engine {
	router := gin.New()
	router.UseRawPath = true
	router.Use(helper.ErrorHandler())

	route.CustomerRoute(router, db, validate)
	route.LoanLimitRoute(router, db, validate)
	route.TransactionRoute(router, db, validate)
	return router
}

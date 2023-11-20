package route

import (
	"xyz-multifinance/controller"
	"xyz-multifinance/repository"
	"xyz-multifinance/service"

	"xyz-multifinance/helper"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func TransactionRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {

	transactionService := service.NewTransactionService(
		repository.NewTransactionRepository(),
		db,
		validate,
	)
	transactionController := controller.NewTransactionController(transactionService)

	router.GET("/transactions", helper.Auth(transactionController.FindAll, []string{helper.RoleAdministrator}))
	router.POST("/transactions", helper.Auth(transactionController.Create, []string{helper.RoleAdministrator}))
	router.PUT("/transactions/:transaction_no", helper.Auth(transactionController.Update, []string{helper.RoleAdministrator}))
	router.DELETE("/transactions/:transaction_no", helper.Auth(transactionController.Delete, []string{helper.RoleAdministrator}))
}

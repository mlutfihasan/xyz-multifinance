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

func LoanLimitRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {

	loanLimitService := service.NewLoanLimitService(
		repository.NewLoanLimitRepository(),
		db,
		validate,
	)
	loanLimitController := controller.NewLoanLimitController(loanLimitService)

	router.GET("/loan_limits", helper.Auth(loanLimitController.FindAll, []string{helper.RoleAdministrator}))
	router.POST("/loan_limits", helper.Auth(loanLimitController.Create, []string{helper.RoleAdministrator}))
	router.PUT("/loan_limits/:loan_limit_id", helper.Auth(loanLimitController.Update, []string{helper.RoleAdministrator}))
	router.DELETE("/loan_limits/:loan_limit_id", helper.Auth(loanLimitController.Delete, []string{helper.RoleAdministrator}))
}

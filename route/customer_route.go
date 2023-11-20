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

func CustomerRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {

	customerService := service.NewCustomerService(
		repository.NewCustomerRepository(),
		db,
		validate,
	)
	customerController := controller.NewCustomerController(customerService)

	router.GET("/customers", helper.Auth(customerController.FindAll, []string{helper.RoleAdministrator}))
	router.POST("/customers", helper.Auth(customerController.Create, []string{helper.RoleAdministrator}))
	router.PUT("/customers/:customer_nik", helper.Auth(customerController.Update, []string{helper.RoleAdministrator}))
	router.DELETE("/customers/:customer_nik", helper.Auth(customerController.Delete, []string{helper.RoleAdministrator}))
}

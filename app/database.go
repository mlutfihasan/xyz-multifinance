package app

import (
	"xyz-multifinance/helper"
	"xyz-multifinance/model/domain"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/gorm"
)

func ConnectDatabase(configuration helper.Configuration) *gorm.DB {
	database := helper.ConnectMysqlDatabaseResolver(configuration)

	if err := database.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}

	err := database.AutoMigrate(
		&domain.Customer{},
		&domain.LoanLimit{},
		&domain.Transaction{},
	)
	if err != nil {
		panic("failed to auto migrate schema")
	}

	return database
}

package repository

import (
	"xyz-multifinance/model/domain"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindAll(db *gorm.DB, filters *map[string]string) domain.Customers
	Create(db *gorm.DB, customer *domain.Customer) *domain.Customer
	Update(db *gorm.DB, customerNik *string, customer *domain.Customer) *domain.Customer
	Delete(db *gorm.DB, customerNik *string)
}

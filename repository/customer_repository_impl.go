package repository

import (
	"xyz-multifinance/helper"
	"xyz-multifinance/model/domain"

	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepositoryImpl{}
}

func (repository *CustomerRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.Customers {
	customers := domain.Customers{}
	tx := db.Model(&domain.Customer{})

	err := helper.ApplyFilter(tx, filters)
	helper.PanicIfError(err)

	err = tx.Find(&customers).Error
	helper.PanicIfError(err)

	return customers
}

func (repository *CustomerRepositoryImpl) Create(db *gorm.DB, customer *domain.Customer) *domain.Customer {
	err := db.Create(&customer).Error
	helper.PanicIfError(err)
	return customer
}

func (repository *CustomerRepositoryImpl) Update(db *gorm.DB, customerNik *string, customer *domain.Customer) *domain.Customer {
	err := db.Model(&domain.Customer{}).
		Where(&domain.Customer{
			CustomerNik: *customerNik,
		}).
		Updates(&customer).Error
	helper.PanicIfError(err)

	err = db.First(&customer, &domain.Customer{CustomerNik: *customerNik}).Error
	helper.PanicIfError(err)

	return customer
}

func (repository *CustomerRepositoryImpl) Delete(db *gorm.DB, customerNik *string) {
	customer := &domain.Customer{}
	tx := db.First(customer, &domain.Customer{CustomerNik: *customerNik}).Updates(&domain.Customer{
		CustomerNik: *customerNik,
	})

	// Deleting the Customer from the database.
	err := tx.Unscoped().Delete(customer, &domain.Customer{CustomerNik: *customerNik}).Error
	helper.PanicIfError(err)
}

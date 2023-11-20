package domain

import (
	"time"

	"xyz-multifinance/model/web"
)

type Customers []Customer
type Customer struct {
	// Fields
	CustomerNik       string    `gorm:"primarykey;size:16;not null;"`
	CustomerName      string    `gorm:"size:60;not null;"`
	CustomerLegalName string    `gorm:"size:60;not null;"`
	PlaceOfBirth      string    `gorm:"size:30;not null;"`
	DateOfBirth       time.Time `gorm:"not null;"`
	Salary            float32   `gorm:"not null;"`
	IdPhoto           string    `gorm:"size:50;not null;"`
	SelfiePhoto       string    `gorm:"size:50;not null;"`

	// Foreign Key Fields
	LoanLimits   []LoanLimit   `gorm:"foreignKey:CustomerNik;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Transactions []Transaction `gorm:"foreignKey:CustomerNik;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

func (customer *Customer) ToCustomerResponse() web.CustomerResponse {
	return web.CustomerResponse{
		// Fields
		CustomerNik:       customer.CustomerNik,
		CustomerName:      customer.CustomerName,
		CustomerLegalName: customer.CustomerLegalName,
		PlaceOfBirth:      customer.PlaceOfBirth,
		DateOfBirth:       customer.DateOfBirth,
		Salary:            customer.Salary,
		IdPhoto:           customer.IdPhoto,
		SelfiePhoto:       customer.SelfiePhoto,
	}
}

func (customers Customers) ToCustomerResponses() []web.CustomerResponse {
	customerResponses := []web.CustomerResponse{}
	for _, customer := range customers {
		customerResponses = append(customerResponses, customer.ToCustomerResponse())
	}
	return customerResponses
}

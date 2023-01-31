package models

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	CustomerNik       string    `json:"customer_nik"`
	CustomerName      string    `json:"customer_name"`
	CustomerLegalName string    `json:"customer_legal_name"`
	PlaceOfBirth      string    `json:"place_of_birth"`
	DateOfBirth       time.Time `json:"date_of_birth"`
	Salary            float32   `json:"salary"`
	IdPhoto           string    `json:"id_photo"`
	SelfiePhoto       string    `json:"selfie_photo"`
}

func (c *Customer) Validate() error {
	if c.CustomerNik == "" {
		return errors.New("Required NIK")
	}

	if c.CustomerName == "" {
		return errors.New("Required Name")
	}

	if c.CustomerLegalName == "" {
		return errors.New("Required Legal Name")
	}

	if c.PlaceOfBirth == "" {
		return errors.New("Required Username")
	}

	if c.Salary == 0 {
		return errors.New("Required Salary")
	}

	if c.IdPhoto == "" {
		return errors.New("Required ID Photo")
	}

	if c.SelfiePhoto == "" {
		return errors.New("Required Selfie Photo")
	}

	return nil
}

func (c *Customer) SaveCustomer(db *gorm.DB) CrudResult {
	err := db.Debug().Create(&c).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			err = errors.New("Data Already Exist")
		}

		return CrudResult{
			Status: "0",
			Note:   err,
		}
	}

	return CrudResult{
		Status: "1",
		Note:   nil,
	}
}

func FindCustomer(db *gorm.DB, nik string) (*[]Customer, error) {
	var custs []Customer

	tx := db.Debug()

	if nik != "" {
		tx = tx.Where(&Customer{CustomerNik: nik})
	}

	err := tx.Find(&custs).Error
	if err != nil {
		return &[]Customer{}, err
	}

	return &custs, nil
}

func UpdateCustomer(db *gorm.DB, customerNikEdit string, updatedColoumn Customer) CrudResult {
	var c Customer

	err := db.Debug().Where(&Customer{CustomerNik: customerNikEdit}).Take(&c).Error
	if err != nil {
		return CrudResult{
			Status: "0",
			Note:   err,
		}
	}

	err = BeforeUpdateCustomer(db, c)
	if err != nil {
		return CrudResult{
			Status: "0",
			Note:   err,
		}
	}

	err = db.Debug().Where(&Customer{CustomerNik: customerNikEdit}).Take(&Customer{}).Updates(updatedColoumn).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			err = errors.New("Data Already Exist")
		}

		return CrudResult{
			Status: "0",
			Note:   err,
		}
	}

	return CrudResult{
		Status: "1",
		Note:   nil,
	}
}

func DeleteCustomer(db *gorm.DB, customerNikDelete string) CrudResult {
	var c Customer

	err := db.Debug().Where(&Customer{CustomerNik: customerNikDelete}).Take(&c).Error
	if err != nil {
		return CrudResult{
			Status: "0",
			Note:   err,
		}
	}

	err = BeforeDeleteCustomer(db, c)
	if err != nil {
		return CrudResult{
			Status: "0",
			Note:   err,
		}
	}

	err = db.Debug().Where(&Customer{CustomerNik: customerNikDelete}).Take(&Customer{}).Delete(&Customer{}).Error
	if err != nil {
		return CrudResult{
			Status: "0",
			Note:   err,
		}
	}

	return CrudResult{
		Status: "1",
		Note:   nil,
	}
}

func BeforeUpdateCustomer(db *gorm.DB, data Customer) error {
	hist := CustomerHistory{}
	hist.Prepare("UPDATE", data)

	err := SaveCustomerHistory(db, hist)
	return err
}

func BeforeDeleteCustomer(db *gorm.DB, data Customer) error {
	hist := CustomerHistory{}
	hist.Prepare("DELETE", data)

	err := SaveCustomerHistory(db, hist)
	return err
}

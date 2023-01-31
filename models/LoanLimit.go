package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type LoanLimit struct {
	LoanId       int     `json:"loan_id"`
	CustomerNik  string  `json:"customer_nik"`
	CustomerName string  `json:"customer_name"`
	OneMonth     float32 `json:"one_month"`
	TwoMonth     float32 `json:"two_month"`
	ThreeMonth   float32 `json:"three_month"`
	FourMonth    float32 `json:"four_month"`
}

func (ll *LoanLimit) Validate() error {
	if ll.CustomerNik == "" {
		return errors.New("Required NIK")
	}

	if ll.CustomerName == "" {
		return errors.New("Required Name")
	}

	if ll.OneMonth == 0 {
		return errors.New("Required One Month Limit")
	}

	if ll.TwoMonth == 0 {
		return errors.New("Required Two Month Limit")
	}

	if ll.ThreeMonth == 0 {
		return errors.New("Required Three Month Limit")
	}

	if ll.FourMonth == 0 {
		return errors.New("Required Four Month Limit")
	}

	return nil
}

func (ll *LoanLimit) SaveLoanLimit(db *gorm.DB) CrudResult {
	err := db.Debug().Create(&ll).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			err = errors.New("Data Already Exist")
		}

		if strings.Contains(err.Error(), "Cannot add or update a child row") {
			err = errors.New("Data Customer Not Exist")
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

func FindLoanLimit(db *gorm.DB, loanId int) (*[]LoanLimit, error) {
	var loanLimits []LoanLimit

	tx := db.Debug()

	if loanId != 0 {
		tx = tx.Where(&LoanLimit{LoanId: loanId})
	}

	err := tx.Find(&loanLimits).Error
	if err != nil {
		return &[]LoanLimit{}, err
	}

	return &loanLimits, nil
}

func UpdateLoanLimit(db *gorm.DB, loanIdEdit int, updatedColoumn LoanLimit) CrudResult {
	var ll LoanLimit

	err := db.Debug().Where(&LoanLimit{LoanId: loanIdEdit}).Take(&ll).Error
	if err != nil {
		return CrudResult{
			Status: "0",
			Note:   err,
		}
	}

	err = BeforeUpdateLoanLimit(db, ll)
	if err != nil {
		return CrudResult{
			Status: "0",
			Note:   err,
		}
	}

	err = db.Debug().Where(&LoanLimit{LoanId: loanIdEdit}).Take(&LoanLimit{}).Updates(updatedColoumn).Error
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

func DeleteLoanLimit(db *gorm.DB, loanIdDelete int) CrudResult {
	var ll LoanLimit

	err := db.Debug().Where(&LoanLimit{LoanId: loanIdDelete}).Take(&ll).Error
	if err != nil {
		return CrudResult{
			Status: "0",
			Note:   err,
		}
	}

	err = BeforeDeleteLoanLimit(db, ll)
	if err != nil {
		return CrudResult{
			Status: "0",
			Note:   err,
		}
	}

	err = db.Debug().Where(&LoanLimit{LoanId: loanIdDelete}).Take(&LoanLimit{}).Delete(&LoanLimit{}).Error
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

func BeforeUpdateLoanLimit(db *gorm.DB, data LoanLimit) error {
	hist := LoanLimitHistory{}
	hist.Prepare("UPDATE", data)

	err := SaveLoanLimitHistory(db, hist)
	return err
}

func BeforeDeleteLoanLimit(db *gorm.DB, data LoanLimit) error {
	hist := LoanLimitHistory{}
	hist.Prepare("DELETE", data)

	err := SaveLoanLimitHistory(db, hist)
	return err
}

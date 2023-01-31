package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	TransactionNo   string    `json:"transaction_no"`
	TransactionDate time.Time `json:"transaction_date"`
	CustomerNik     string    `json:"customer_nik"`
	CustomerName    string    `json:"customer_name"`
	OnTheRoad       float32   `json:"on_the_road"`
	AdminFee        float32   `json:"admin_fee"`
	LoanAmount      float32   `json:"loan_amount"`
	InterestAmount  float32   `json:"interest_amount"`
	AssetName       string    `json:"asset_name"`
}

func (t *Transaction) Prepare() {
	t.TransactionDate = time.Now()

	unixTransactionDate := strconv.Itoa(int(t.TransactionDate.Unix()))
	t.TransactionNo = fmt.Sprintf("%s/%s", t.CustomerNik, unixTransactionDate)
}

func (t *Transaction) Validate() error {
	if t.TransactionNo == "" {
		return errors.New("Required Transaction Number")
	}

	if t.CustomerNik == "" {
		return errors.New("Required NIK")
	}

	if t.CustomerName == "" {
		return errors.New("Required Name")
	}

	if t.OnTheRoad == 0 {
		return errors.New("Required On The Road")
	}

	if t.AdminFee == 0 {
		return errors.New("Required Admin Fee")
	}

	if t.LoanAmount == 0 {
		return errors.New("Required Loan Amount")
	}

	if t.InterestAmount == 0 {
		return errors.New("Required Interest Amount")
	}

	if t.AssetName == "" {
		return errors.New("Required Asset Name")
	}

	return nil
}

func (t *Transaction) SaveTransaction(db *gorm.DB) CrudResult {
	var ll LoanLimit
	err := db.Debug().Where(&LoanLimit{CustomerNik: t.CustomerNik}).Take(&ll).Error
	if err != nil {
		return CrudResult{
			Status: "0",
			Note:   err,
		}
	}

	var tb []Transaction
	db.Debug().Where(&Transaction{CustomerNik: t.CustomerNik}).Take(&tb)

	var totalTransaction float32
	for _, data := range tb {
		totalTransaction = totalTransaction + data.OnTheRoad
	}

	totalTransaction = totalTransaction + t.OnTheRoad

	if totalTransaction > ll.OneMonth {
		if totalTransaction > ll.TwoMonth {
			if totalTransaction > ll.ThreeMonth {
				if totalTransaction > ll.FourMonth {
					err = errors.New("Exceeds The Limit")

					return CrudResult{
						Status: "0",
						Note:   err,
					}
				}
			}
		}
	}

	err = db.Debug().Create(&t).Error
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

func FindTransaction(db *gorm.DB, transactionNo string) (*[]Transaction, error) {
	var transactions []Transaction

	tx := db.Debug()

	if transactionNo != "" {
		tx = tx.Where(&Transaction{TransactionNo: transactionNo})
	}

	err := tx.Find(&transactions).Error
	if err != nil {
		return &[]Transaction{}, err
	}

	return &transactions, nil
}

func DeleteTransaction(db *gorm.DB, transactionNoDelete string) CrudResult {
	var t Transaction

	err := db.Debug().Where(&Transaction{TransactionNo: transactionNoDelete}).Take(&t).Error
	if err != nil {
		return CrudResult{
			Status: "0",
			Note:   err,
		}
	}

	err = BeforeDeleteTransaction(db, t)
	if err != nil {
		return CrudResult{
			Status: "0",
			Note:   err,
		}
	}

	err = db.Debug().Where(&Transaction{TransactionNo: transactionNoDelete}).Take(&Transaction{}).Delete(&Transaction{}).Error
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

func BeforeDeleteTransaction(db *gorm.DB, data Transaction) error {
	hist := TransactionHistory{}
	hist.Prepare("DELETE", data)

	err := SaveTransactionHistory(db, hist)
	return err
}

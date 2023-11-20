package test

import (
	"log"
	"testing"
	"xyz-multifinance/app"
	"xyz-multifinance/helper"
	"xyz-multifinance/model/domain"
	"xyz-multifinance/repository"

	"github.com/stretchr/testify/require"
)

var testLoanLimit *repository.LoanLimitRepositoryImpl

func TestCreateLoanLimit(t *testing.T) {
	loanLimit := domain.LoanLimit{
		CustomerNik:  "2893876789876542",
		CustomerName: "John Doe",
		OneMonth:     10000000.00,
		TwoMonth:     1000000.00,
		ThreeMonth:   100000.00,
		FourMonth:    10000.00,
	}

	configuration, err := helper.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db := app.ConnectDatabase(configuration)

	loanLimitData := testLoanLimit.Create(db, &loanLimit)
	require.NotEmpty(t, loanLimitData)

	require.Equal(t, loanLimit.LoanLimitID, loanLimitData.LoanLimitID)
	require.Equal(t, loanLimit.CustomerNik, loanLimitData.CustomerNik)
	require.Equal(t, loanLimit.CustomerName, loanLimitData.CustomerName)
	require.Equal(t, loanLimit.OneMonth, loanLimitData.OneMonth)
	require.Equal(t, loanLimit.TwoMonth, loanLimitData.TwoMonth)
	require.Equal(t, loanLimit.ThreeMonth, loanLimitData.ThreeMonth)
	require.Equal(t, loanLimit.FourMonth, loanLimitData.FourMonth)

	require.NotZero(t, loanLimitData.LoanLimitID)
	require.NotZero(t, loanLimitData.CustomerNik)
	require.NotZero(t, loanLimitData.CustomerName)
	require.NotZero(t, loanLimitData.OneMonth)
	require.NotZero(t, loanLimitData.TwoMonth)
	require.NotZero(t, loanLimitData.ThreeMonth)
	require.NotZero(t, loanLimitData.FourMonth)
}

func TestGetLoanLimit(t *testing.T) {
	configuration, err := helper.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db := app.ConnectDatabase(configuration)

	var filters map[string]string
	loanLimitData := testLoanLimit.FindAll(db, &filters)
	require.NotEmpty(t, loanLimitData)
}

func TestUpdateLoanLimit(t *testing.T) {
	configuration, err := helper.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db := app.ConnectDatabase(configuration)

	loanLimitID := "1"
	loanLimit := domain.LoanLimit{
		OneMonth:   20000000.00,
		TwoMonth:   2000000.00,
		ThreeMonth: 200000.00,
		FourMonth:  20000.00,
	}

	loanLimitData := testLoanLimit.Update(db, &loanLimitID, &loanLimit)
	require.NotEmpty(t, loanLimitData)
}

func TestDeleteLoanLimit(t *testing.T) {
	configuration, err := helper.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db := app.ConnectDatabase(configuration)

	loanLimitID := "1"

	loanLimitData := testLoanLimit.Delete(db, &loanLimitID)
	require.NotEmpty(t, loanLimitData)
}

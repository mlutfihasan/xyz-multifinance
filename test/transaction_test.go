package test

import (
	"log"
	"strconv"
	"testing"
	"time"
	"xyz-multifinance/app"
	"xyz-multifinance/helper"
	"xyz-multifinance/model/domain"
	"xyz-multifinance/repository"

	"github.com/stretchr/testify/require"
)

var testTransaction *repository.TransactionRepositoryImpl

func TestCreateTransaction(t *testing.T) {
	unixTimeNow := strconv.Itoa(int(time.Now().Unix()))
	transactionNo := "2893876789876542" + "/" + unixTimeNow

	transaction := domain.Transaction{
		TransactionNo:   transactionNo,
		TransactionDate: time.Now(),
		CustomerNik:     "2893876789876542",
		CustomerName:    "John Doe",
		OnTheRoad:       110000.00,
		AdminFee:        11000.00,
		LoanAmount:      1100000.00,
		InterestAmount:  110000.00,
		AssetName:       "ASET BERHARGA",
	}

	configuration, err := helper.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db := app.ConnectDatabase(configuration)

	transactionData := testTransaction.Create(db, &transaction)
	require.NotEmpty(t, transactionData)

	require.Equal(t, transaction.TransactionNo, transactionData.TransactionNo)
	require.Equal(t, transaction.TransactionDate, transactionData.TransactionDate)
	require.Equal(t, transaction.CustomerNik, transactionData.CustomerNik)
	require.Equal(t, transaction.CustomerName, transactionData.CustomerName)
	require.Equal(t, transaction.OnTheRoad, transactionData.OnTheRoad)
	require.Equal(t, transaction.AdminFee, transactionData.AdminFee)
	require.Equal(t, transaction.LoanAmount, transactionData.LoanAmount)
	require.Equal(t, transaction.InterestAmount, transactionData.InterestAmount)
	require.Equal(t, transaction.AssetName, transactionData.AssetName)

	require.NotZero(t, transactionData.TransactionNo)
	require.NotZero(t, transactionData.TransactionDate)
	require.NotZero(t, transactionData.CustomerNik)
	require.NotZero(t, transactionData.CustomerName)
	require.NotZero(t, transactionData.OnTheRoad)
	require.NotZero(t, transactionData.AdminFee)
	require.NotZero(t, transactionData.LoanAmount)
	require.NotZero(t, transactionData.InterestAmount)
	require.NotZero(t, transactionData.AssetName)
}

func TestGetTransaction(t *testing.T) {
	configuration, err := helper.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db := app.ConnectDatabase(configuration)

	var filters map[string]string
	transactionData := testTransaction.FindAll(db, &filters)
	require.NotEmpty(t, transactionData)
}

func TestUpdateTransaction(t *testing.T) {
	configuration, err := helper.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db := app.ConnectDatabase(configuration)

	transactionID := "2893876789876542/1700491342"
	transaction := domain.Transaction{
		OnTheRoad:      220000.00,
		AdminFee:       22000.00,
		LoanAmount:     2200000.00,
		InterestAmount: 220000.00,
		AssetName:      "ASET EMAS",
	}

	transactionData := testTransaction.Update(db, &transactionID, &transaction)
	require.NotEmpty(t, transactionData)
}

func TestDeleteTransaction(t *testing.T) {
	configuration, err := helper.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db := app.ConnectDatabase(configuration)

	transactionID := "2893876789876542/1700491342"

	transactionData := testTransaction.Delete(db, &transactionID)
	require.NotEmpty(t, transactionData)
}

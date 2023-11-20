package test

import (
	"fmt"
	"log"
	"testing"
	"time"
	"xyz-multifinance/app"
	"xyz-multifinance/helper"
	"xyz-multifinance/model/domain"
	"xyz-multifinance/repository"

	"github.com/stretchr/testify/require"
)

var testCustomer *repository.CustomerRepositoryImpl

func TestCreateCustomer(t *testing.T) {
	layoutDate := "2006-01-02 15:04:05"
	useDate := "1991-01-01 00:00:00"
	dateOfBirth, err := time.Parse(layoutDate, useDate)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	customer := domain.Customer{
		CustomerNik:       "2893876789876542",
		CustomerName:      "John Doe",
		CustomerLegalName: "John Doe OConnor",
		PlaceOfBirth:      "Las Vegas",
		DateOfBirth:       dateOfBirth,
		Salary:            11000000.00,
		IdPhoto:           "./file/ID/2893876789876542-ID.jpg",
		SelfiePhoto:       "./file/SELFIE/2893876789876542-SELFIE.jpg",
	}

	configuration, err := helper.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db := app.ConnectDatabase(configuration)

	customerData := testCustomer.Create(db, &customer)
	require.NotEmpty(t, customerData)

	require.Equal(t, customer.CustomerNik, customerData.CustomerNik)
	require.Equal(t, customer.CustomerName, customerData.CustomerName)
	require.Equal(t, customer.CustomerLegalName, customerData.CustomerLegalName)
	require.Equal(t, customer.PlaceOfBirth, customerData.PlaceOfBirth)
	require.Equal(t, customer.DateOfBirth, customerData.DateOfBirth)
	require.Equal(t, customer.Salary, customerData.Salary)
	require.Equal(t, customer.IdPhoto, customerData.IdPhoto)
	require.Equal(t, customer.SelfiePhoto, customerData.SelfiePhoto)

	require.NotZero(t, customerData.CustomerNik)
	require.NotZero(t, customerData.CustomerName)
	require.NotZero(t, customerData.CustomerLegalName)
	require.NotZero(t, customerData.PlaceOfBirth)
	require.NotZero(t, customerData.DateOfBirth)
	require.NotZero(t, customerData.Salary)
	require.NotZero(t, customerData.IdPhoto)
	require.NotZero(t, customerData.SelfiePhoto)
}

func TestGetCustomer(t *testing.T) {
	configuration, err := helper.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db := app.ConnectDatabase(configuration)

	var filters map[string]string
	customerData := testCustomer.FindAll(db, &filters)
	require.NotEmpty(t, customerData)
}

func TestUpdateCustomer(t *testing.T) {
	configuration, err := helper.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db := app.ConnectDatabase(configuration)

	customerNik := "2893876789876542"
	customer := domain.Customer{
		Salary: 22000000.00,
	}

	customerData := testCustomer.Update(db, &customerNik, &customer)
	require.NotEmpty(t, customerData)
}

func TestDeleteCustomer(t *testing.T) {
	configuration, err := helper.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db := app.ConnectDatabase(configuration)

	customerNik := "2893876789876542"

	customerData := testCustomer.Delete(db, &customerNik)
	require.NotEmpty(t, customerData)
}

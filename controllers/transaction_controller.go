package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"xyz-multifinance/models"
	"xyz-multifinance/responses"
)

func (server *Server) TransactionController(w http.ResponseWriter, r *http.Request) {
	transaction := models.Transaction{}
	vars := r.URL.Query()

	if r.Method == "POST" {
		var err error

		customerNik := vars.Get("customer_nik")
		customerName := vars.Get("customer_name")
		onTheRoad := vars.Get("on_the_road")
		adminFee := vars.Get("admin_fee")
		loanAmount := vars.Get("loan_amount")
		interestAmount := vars.Get("interest_amount")
		assetName := vars.Get("asset_name")

		transaction.CustomerNik = customerNik
		transaction.CustomerName = customerName
		transaction.AssetName = assetName

		if onTheRoad != "" {
			onTheRoads, err := strconv.ParseFloat(onTheRoad, 32)
			if err != nil {
				responses.ERROR(w, http.StatusOK, err)
				return
			}
			transaction.OnTheRoad = float32(onTheRoads)
		}

		if adminFee != "" {
			adminFees, err := strconv.ParseFloat(adminFee, 32)
			if err != nil {
				responses.ERROR(w, http.StatusOK, err)
				return
			}
			transaction.AdminFee = float32(adminFees)
		}

		if loanAmount != "" {
			loanAmounts, err := strconv.ParseFloat(loanAmount, 32)
			if err != nil {
				responses.ERROR(w, http.StatusOK, err)
				return
			}
			transaction.LoanAmount = float32(loanAmounts)
		}

		if interestAmount != "" {
			interestAmounts, err := strconv.ParseFloat(interestAmount, 32)
			if err != nil {
				responses.ERROR(w, http.StatusOK, err)
				return
			}
			transaction.InterestAmount = float32(interestAmounts)
		}

		transaction.Prepare()
		err = transaction.Validate()
		if err != nil {
			responses.ERROR(w, http.StatusOK, err)
			return
		}

		responseCreated := transaction.SaveTransaction(server.DB)
		if responseCreated.Status == "0" {
			responses.ERROR(w, http.StatusOK, responseCreated.Note)
			return
		}

		responses.JSON(w, http.StatusCreated, responseCreated)
		return
	}

	if r.Method == "GET" {
		transactionNo := vars.Get("transaction_no")

		dataTrans, err := models.FindTransaction(server.DB, transactionNo)
		if err != nil {
			responses.ERROR(w, http.StatusOK, err)
			return
		}

		responses.SUCCESS(w, http.StatusOK, err, dataTrans)
		return
	}

	if r.Method == "DELETE" {
		transactionNoDelete := vars.Get("transaction_no_delete")
		if transactionNoDelete == "" {
			responses.ERROR(w, http.StatusOK, errors.New("Required Transaction Number"))
			return
		}

		responseDeleted := models.DeleteTransaction(server.DB, transactionNoDelete)
		if responseDeleted.Status == "0" {
			responses.ERROR(w, http.StatusOK, responseDeleted.Note)
			return
		}

		responses.JSON(w, http.StatusOK, responseDeleted)
		return
	}
}

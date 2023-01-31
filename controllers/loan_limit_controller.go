package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"xyz-multifinance/models"
	"xyz-multifinance/responses"
)

func (server *Server) LoanLimitController(w http.ResponseWriter, r *http.Request) {
	loanLimit := models.LoanLimit{}
	vars := r.URL.Query()

	if r.Method == "POST" {
		var err error

		customerNik := vars.Get("customer_nik")
		customerName := vars.Get("customer_name")
		oneMonth := vars.Get("one_month")
		twoMonth := vars.Get("two_month")
		threeMonth := vars.Get("three_month")
		fourMonth := vars.Get("four_month")

		loanLimit.CustomerNik = customerNik
		loanLimit.CustomerName = customerName

		if oneMonth != "" {
			oneMonths, err := strconv.ParseFloat(oneMonth, 32)
			if err != nil {
				responses.ERROR(w, http.StatusOK, err)
				return
			}
			loanLimit.OneMonth = float32(oneMonths)
		}

		if twoMonth != "" {
			twoMonths, err := strconv.ParseFloat(twoMonth, 32)
			if err != nil {
				responses.ERROR(w, http.StatusOK, err)
				return
			}
			loanLimit.TwoMonth = float32(twoMonths)
		}

		if threeMonth != "" {
			threeMonths, err := strconv.ParseFloat(threeMonth, 32)
			if err != nil {
				responses.ERROR(w, http.StatusOK, err)
				return
			}
			loanLimit.ThreeMonth = float32(threeMonths)
		}

		if fourMonth != "" {
			fourMonths, err := strconv.ParseFloat(fourMonth, 32)
			if err != nil {
				responses.ERROR(w, http.StatusOK, err)
				return
			}
			loanLimit.FourMonth = float32(fourMonths)
		}

		err = loanLimit.Validate()
		if err != nil {
			responses.ERROR(w, http.StatusOK, err)
			return
		}

		responseCreated := loanLimit.SaveLoanLimit(server.DB)
		if responseCreated.Status == "0" {
			responses.ERROR(w, http.StatusOK, responseCreated.Note)
			return
		}

		responses.JSON(w, http.StatusCreated, responseCreated)
		return
	}

	if r.Method == "GET" {
		var err error
		loanId := vars.Get("loan_id")

		var loanIdFilter int
		if loanId != "" {
			loanIdFilter, err = strconv.Atoi(loanId)
			if err != nil {
				responses.ERROR(w, http.StatusOK, err)
				return
			}
		}

		dataLoan, err := models.FindLoanLimit(server.DB, loanIdFilter)
		if err != nil {
			responses.ERROR(w, http.StatusOK, err)
			return
		}

		responses.SUCCESS(w, http.StatusOK, err, dataLoan)
		return
	}

	if r.Method == "PUT" {
		loanIdEdit := vars.Get("loan_id_edit")
		if loanIdEdit == "" {
			responses.ERROR(w, http.StatusOK, errors.New("Required Loan ID"))
			return
		}

		loanIdEditFilter, err := strconv.Atoi(loanIdEdit)
		if err != nil {
			responses.ERROR(w, http.StatusOK, err)
			return
		}

		oneMonth := vars.Get("one_month")
		twoMonth := vars.Get("two_month")
		threeMonth := vars.Get("three_month")
		fourMonth := vars.Get("four_month")

		if oneMonth == "" && twoMonth == "" && threeMonth == "" && fourMonth == "" {
			responses.ERROR(w, http.StatusOK, errors.New("Nothing To Edit"))
			return
		}

		if oneMonth != "" {
			oneMonths, err := strconv.ParseFloat(oneMonth, 32)
			if err != nil {
				responses.ERROR(w, http.StatusOK, err)
				return
			}
			loanLimit.OneMonth = float32(oneMonths)
		}

		if twoMonth != "" {
			twoMonths, err := strconv.ParseFloat(twoMonth, 32)
			if err != nil {
				responses.ERROR(w, http.StatusOK, err)
				return
			}
			loanLimit.TwoMonth = float32(twoMonths)
		}

		if threeMonth != "" {
			threeMonths, err := strconv.ParseFloat(threeMonth, 32)
			if err != nil {
				responses.ERROR(w, http.StatusOK, err)
				return
			}
			loanLimit.ThreeMonth = float32(threeMonths)
		}

		if fourMonth != "" {
			fourMonths, err := strconv.ParseFloat(fourMonth, 32)
			if err != nil {
				responses.ERROR(w, http.StatusOK, err)
				return
			}
			loanLimit.FourMonth = float32(fourMonths)
		}

		responseUpdated := models.UpdateLoanLimit(server.DB, loanIdEditFilter, loanLimit)
		if responseUpdated.Status == "0" {
			responses.ERROR(w, http.StatusOK, responseUpdated.Note)
			return
		}

		responses.JSON(w, http.StatusOK, responseUpdated)
		return
	}

	if r.Method == "DELETE" {
		loanIdDelete := vars.Get("loan_id_delete")
		if loanIdDelete == "" {
			responses.ERROR(w, http.StatusOK, errors.New("Required Loan ID"))
			return
		}

		loanIdDeleteFilter, err := strconv.Atoi(loanIdDelete)
		if err != nil {
			responses.ERROR(w, http.StatusOK, err)
			return
		}

		responseDeleted := models.DeleteLoanLimit(server.DB, loanIdDeleteFilter)
		if responseDeleted.Status == "0" {
			responses.ERROR(w, http.StatusOK, responseDeleted.Note)
			return
		}

		responses.JSON(w, http.StatusOK, responseDeleted)
		return
	}
}

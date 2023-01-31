package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"xyz-multifinance/models"
	"xyz-multifinance/responses"
	"xyz-multifinance/utils"
)

func (server *Server) CustomerController(w http.ResponseWriter, r *http.Request) {
	customer := models.Customer{}
	vars := r.URL.Query()

	if r.Method == "POST" {
		var err error

		customerNik := vars.Get("customer_nik")
		customerName := vars.Get("customer_name")
		customerLegalName := vars.Get("customer_legal_name")
		placeOfBirth := vars.Get("place_of_birth")
		dateOfBirth := vars.Get("date_of_birth")
		salary := vars.Get("salary")

		customer.CustomerNik = customerNik
		customer.CustomerName = customerName
		customer.CustomerLegalName = customerLegalName
		customer.PlaceOfBirth = placeOfBirth
		customer.DateOfBirth, err = time.Parse("2006-01-02", dateOfBirth)
		if err != nil {
			responses.ERROR(w, http.StatusOK, err)
			return
		}

		if salary != "" {
			salaries, err := strconv.ParseFloat(salary, 32)
			if err != nil {
				responses.ERROR(w, http.StatusOK, err)
				return
			}
			customer.Salary = float32(salaries)
		}

		idPhoto, handlerIdPhoto, _ := r.FormFile("id_photo")
		if idPhoto != nil && handlerIdPhoto != nil {
			linkIdPhoto, err := utils.UploadFilePhoto(idPhoto, *handlerIdPhoto, customer.CustomerNik, "ID PHOTO")
			if err != nil {
				responses.ERROR(w, http.StatusOK, err)
				return
			}

			customer.IdPhoto = linkIdPhoto
		} else {
			responses.ERROR(w, http.StatusOK, errors.New("Required ID Photo"))
			return
		}

		selfiePhoto, handlerSelfiePhoto, _ := r.FormFile("selfie_photo")
		if selfiePhoto != nil && handlerSelfiePhoto != nil {
			linkSelfiePhoto, err := utils.UploadFilePhoto(selfiePhoto, *handlerSelfiePhoto, customer.CustomerNik, "SELFIE PHOTO")
			if err != nil {
				responses.ERROR(w, http.StatusOK, err)
				return
			}

			customer.SelfiePhoto = linkSelfiePhoto
		} else {
			responses.ERROR(w, http.StatusOK, errors.New("Required Selfie Photo"))
			return
		}

		err = customer.Validate()
		if err != nil {
			responses.ERROR(w, http.StatusOK, err)
			return
		}

		responseCreated := customer.SaveCustomer(server.DB)
		if responseCreated.Status == "0" {
			responses.ERROR(w, http.StatusOK, responseCreated.Note)
			return
		}

		responses.JSON(w, http.StatusCreated, responseCreated)
		return
	}

	if r.Method == "GET" {
		customerNik := vars.Get("customer_nik")

		dataCust, err := models.FindCustomer(server.DB, customerNik)
		if err != nil {
			responses.ERROR(w, http.StatusOK, err)
			return
		}

		responses.SUCCESS(w, http.StatusOK, err, dataCust)
		return
	}

	if r.Method == "PUT" {
		customerNikEdit := vars.Get("customer_nik_edit")
		if customerNikEdit == "" {
			responses.ERROR(w, http.StatusOK, errors.New("Required NIK"))
			return
		}

		customerName := vars.Get("customer_name")
		salary := vars.Get("salary")
		if customerName == "" && salary == "" {
			responses.ERROR(w, http.StatusOK, errors.New("Nothing To Edit"))
			return
		}

		customer.CustomerName = customerName
		if salary != "" {
			salaries, err := strconv.ParseFloat(salary, 32)
			if err != nil {
				responses.ERROR(w, http.StatusOK, err)
				return
			}
			customer.Salary = float32(salaries)
		}

		responseUpdated := models.UpdateCustomer(server.DB, customerNikEdit, customer)
		if responseUpdated.Status == "0" {
			responses.ERROR(w, http.StatusOK, responseUpdated.Note)
			return
		}

		responses.JSON(w, http.StatusOK, responseUpdated)
		return
	}

	if r.Method == "DELETE" {
		customerNikDelete := vars.Get("customer_nik_delete")
		if customerNikDelete == "" {
			responses.ERROR(w, http.StatusOK, errors.New("Required NIK"))
			return
		}

		responseDeleted := models.DeleteCustomer(server.DB, customerNikDelete)
		if responseDeleted.Status == "0" {
			responses.ERROR(w, http.StatusOK, responseDeleted.Note)
			return
		}

		responses.JSON(w, http.StatusOK, responseDeleted)
		return
	}
}

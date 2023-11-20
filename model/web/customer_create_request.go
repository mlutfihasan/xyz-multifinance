package web

import (
	"mime/multipart"
	"time"
)

type CustomerCreateRequest struct {
	// Fields
	CustomerNik       string    `json:"customer_nik"`
	CustomerName      string    `json:"customer_name"`
	CustomerLegalName string    `json:"customer_legal_name"`
	PlaceOfBirth      string    `json:"place_of_birth"`
	DateOfBirth       time.Time `json:"date_of_birth"`
	Salary            float32   `json:"salary"`
	IdPhoto           *multipart.FileHeader
	SelfiePhoto       *multipart.FileHeader
}

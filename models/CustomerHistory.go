package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomerHistory struct {
	CustomerNik       string    `json:"customer_nik"`
	CustomerName      string    `json:"customer_name"`
	CustomerLegalName string    `json:"customer_legal_name"`
	PlaceOfBirth      string    `json:"place_of_birth"`
	DateOfBirth       time.Time `json:"date_of_birth"`
	Salary            float32   `json:"salary"`
	IdPhoto           string    `json:"id_photo"`
	SelfiePhoto       string    `json:"selfie_photo"`

	ActionTaken string    `json:"action_taken"`
	ActionTime  time.Time `json:"action_time"`
}

func (ch *CustomerHistory) Prepare(action string, data Customer) {
	ch.CustomerNik = data.CustomerNik
	ch.CustomerName = data.CustomerName
	ch.CustomerLegalName = data.CustomerLegalName
	ch.PlaceOfBirth = data.PlaceOfBirth
	ch.DateOfBirth = data.DateOfBirth
	ch.Salary = data.Salary
	ch.IdPhoto = data.IdPhoto
	ch.SelfiePhoto = data.SelfiePhoto

	ch.ActionTaken = action
	ch.ActionTime = time.Now()
}

func SaveCustomerHistory(db *gorm.DB, data CustomerHistory) error {
	err := db.Debug().Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}

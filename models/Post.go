package models

type CrudResult struct {
	Status string `json:"status"`
	Note   error  `json:"note"`
}

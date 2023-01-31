package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {

		errorMessage := err.Error()

		JSON(w, statusCode, struct {
			Status string `json:"status"`
			Error  string `json:"note"`
		}{
			Status: "0",
			Error:  errorMessage,
		})
		return
	}

	JSON(w, http.StatusOK, nil)
}

func SUCCESS(w http.ResponseWriter, statusCode int, err error, data interface{}) {
	if err != nil {
		JSON(w, statusCode, struct {
			Status string `json:"status"`
			Error  string `json:"note"`
		}{
			Status: "1",
			Error:  err.Error(),
		})
		return
	} else {
		JSON(w, statusCode, struct {
			Status string      `json:"status"`
			Items  interface{} `json:"data"`
			Error  string      `json:"note"`
		}{
			Status: "1",
			Items:  data,
			Error:  "",
		})
		return
	}
}

package main

import (
	"log"
	"net/http"

	"xyz-multifinance/app"
	"xyz-multifinance/helper"

	"github.com/go-playground/validator/v10"
)

func main() {
	configuration, err := helper.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	port := configuration.Port
	db := app.ConnectDatabase(configuration)

	// Validator
	validate := validator.New()

	router := app.NewRouter(db, validate)
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Printf("Server is running on port %s", port)

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}

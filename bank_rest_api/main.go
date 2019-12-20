package main

import (
	"GO-examples/bank_rest_api/controllers"
	"GO-examples/bank_rest_api/repositories"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "draqneel_db"
)

func main(){
	router := mux.NewRouter()
	repository := repositories.NewBankRepo(dbUser, dbPassword, dbName)
	controller := controllers.GetNewDataController(repository)

	router.HandleFunc("/clients", controller.GetAllUsers).Methods("GET")
	//router.HandleFunc("/client/{id}", repository.GetUser).Methods("GET")
	router.HandleFunc("/client", repository.CreateUser).Methods("POST")


	log.Fatal(http.ListenAndServe(":8000", router))
}
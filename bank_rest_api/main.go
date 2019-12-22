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

	router.HandleFunc("/clients", controller.GetAllUsersController).Methods("GET")
	router.HandleFunc("/client", controller.CreateUserController).Methods("POST")
	router.HandleFunc("/client/{id}", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/client/{id}", controller.DeleteUser).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":8000", router))
}
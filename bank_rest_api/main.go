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

	router.HandleFunc("/users", controller.GetAllUsersController).Methods("GET")
	router.HandleFunc("/user", controller.CreateUserController).Methods("POST")
	router.HandleFunc("/user/{id}", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", controller.DeleteUser).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":8000", router))
}
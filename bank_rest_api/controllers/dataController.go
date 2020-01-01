package controllers

import (
	"GO-examples/bank_rest_api/models"
	"GO-examples/bank_rest_api/repositories"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type DataController struct {
	repo repositories.BankRepo
}

func  GetNewDataController(repo *repositories.BankRepo) *DataController {
	return &DataController{repo:*repo}
}

func (controller *DataController) GetAllUsersController(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "application/json")
	users := controller.repo.GetUsersList()
	_ = json.NewEncoder(writer).Encode(users)
}

func (controller *DataController) CreateUserController(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "application/json")
	var user models.User
	_ = json.NewDecoder(request.Body).Decode(&user)
	controller.repo.CreateUser(&user)

	_ = json.NewEncoder(writer).Encode(user)
}

func (controller *DataController) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	var newUser models.User
	_ = json.NewDecoder(request.Body).Decode(&newUser)
	oldUserId, err :=  strconv.Atoi(params["id"])

	if err != nil {
		log.Fatal(err)
	}
	controller.repo.UpdateUser(&newUser, oldUserId)

}

func (controller *DataController) DeleteUser(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	deleteUserId, err :=  strconv.Atoi(params["id"])

	if err != nil {
		log.Fatal(err)
	}

	controller.repo.DeleteUserById(deleteUserId)
}

func (controller *DataController) GetAllEmployees(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "application/json")
	employees := controller.repo.GetEmployeesList()
	_ = json.NewEncoder(writer).Encode(employees)
}

func (controller *DataController) CreateEmployeeController(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "application/json")
	var employee models.Employee

	_ = json.NewDecoder(request.Body).Decode(&employee)
	controller.repo.CreateEmployee(&employee)

	_ = json.NewEncoder(writer).Encode(employee)
}


package controllers

import (
	"GO-examples/bank_rest_api/repositories"
	"encoding/json"
	"net/http"
)

type DataController struct {
	repo repositories.BankRepo
}

func  GetNewDataController(repo *repositories.BankRepo) *DataController {
	return &DataController{repo:*repo}
}

func (controller *DataController) GetAllUsers(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "application/json")
	users := controller.repo.GetUsersList()
	_ = json.NewEncoder(writer).Encode(users)
}


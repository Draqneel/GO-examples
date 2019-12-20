package repositories

import (
	"GO-examples/bank_rest_api/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	//"log"
	"net/http"
	"strconv"
	//"encoding/json"
	//"github.com/gorilla/mux"
)

type BankRepo struct {
	users        []models.User
	employees    []models.Employee
	clients      []models.Client
	branches     []models.Branch
	balances     []models.Balance
	transactions []models.Transaction
	db           *sql.DB
}

func NewBankRepo(dbUser string, dbPassword string, dbName string) *BankRepo {
	connectionStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", connectionStr)

	if err != nil {
		log.Fatal(err)
	}

	return &BankRepo{db: db}
}

// Work with users

func (repo *BankRepo) GetUsersList() []models.User {
	result, err := repo.db.Query("SELECT * FROM users")

	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		var user models.User
		err := result.Scan(&user.ID, &user.Login, &user.HashPassword, &user.FullName, &user.PhoneNumber, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		repo.users = append(repo.users, user)
	}
	result.Close()
	return repo.users
}



//func (repo *BankRepo) GetUser(writer http.ResponseWriter, request *http.Request){
//	writer.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(request)
//
//	_ = json.NewEncoder(writer).Encode(&models.User{})
//}

func (repo *BankRepo) CreateUser(writer http.ResponseWriter, request *http.Request){
	fmt.Print("Пытаюсь создать юзера")
	writer.Header().Set("Content-Type", "application/json")
	var user models.User
	_ = json.NewDecoder(request.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(1000000))
	repo.users = append(repo.users, user)
	fmt.Print(user)
	_ = json.NewEncoder(writer).Encode(user)
}

func (repo *BankRepo) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range repo.users {
		if item.ID == params["id"] {
			repo.users = append(repo.users[:index], repo.users[index+1:]...)
			var user models.User
			_ = json.NewDecoder(request.Body).Decode(&user)
			user.ID = params["id"]
			repo.users = append(repo.users, user)
			_ = json.NewEncoder(writer).Encode(user)
			return
		}
	}
	_ = json.NewEncoder(writer).Encode(repo.users)
}

func (repo *BankRepo) DeleteUser(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for i, item := range repo.users {
		if item.ID == params["id"] {
			// *My deleting algorithm, it must be faster*
			//fmt.Print("before:",repo.users)
			//repo.users[i] = repo.users[len(repo.users)-1]
			//repo.users[len(repo.users)-1] = models.User{}
			//repo.users = repo.users[:len(repo.users) - 1]
			//fmt.Print("after:",repo.users)
			repo.users = append(repo.users[:i], repo.users[i+1:]...)
			break
		}
	}
}
// Work with employees

func (repo *BankRepo) getAllEmployees(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(repo.employees)
}

func (repo *BankRepo) getEmployee(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range repo.employees {
		if item.ID == params["id"] {
			_ = json.NewEncoder(writer).Encode(item)
			return
		}
	}
	_ = json.NewEncoder(writer).Encode(&models.Employee{})
}
// Work with clients

func (repo *BankRepo) getAllClients(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(repo.clients)
}

func (repo *BankRepo) getClient(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range repo.clients {
		if item.ID == params["id"] {
			_ = json.NewEncoder(writer).Encode(item)
			return
		}
	}
	_ = json.NewEncoder(writer).Encode(&models.Client{})
}
// Work with branches

func (repo *BankRepo) getAllBranches(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(repo.branches)
}

func (repo *BankRepo) getBranch(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range repo.branches {
		if item.ID == params["id"] {
			_ = json.NewEncoder(writer).Encode(item)
			return
		}
	}
	_ = json.NewEncoder(writer).Encode(&models.Branch{})
}
// Work with balances

func (repo *BankRepo) getAllBalances(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(repo.balances)
}

func (repo *BankRepo) getBalance(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range repo.balances {
		if item.ID == params["id"] {
			_ = json.NewEncoder(writer).Encode(item)
			return
		}
	}
	_ = json.NewEncoder(writer).Encode(&models.Balance{})
}
// Work with transactions

func (repo *BankRepo) getAllTransactions(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(repo.transactions)
}

func (repo *BankRepo) getTransaction(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range repo.transactions {
		if item.ID == params["id"] {
			_ = json.NewEncoder(writer).Encode(item)
			return
		}
	}
	_ = json.NewEncoder(writer).Encode(&models.Transaction{})
}
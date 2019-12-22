package repositories

import (
	"GO-examples/bank_rest_api/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	//"log"
	"net/http"
	//"encoding/json"
	//"github.com/gorilla/mux"
)

type BankRepo struct {
	users        		[]models.User
	usersFlag	 		bool
	employees    		[]models.Employee
	employeesFlag		bool
	clients      		[]models.Client
	clientsFlag			bool
	branches     		[]models.Branch
	branchesFlag		bool
	balances     		[]models.Balance
	balancesFlag		bool
	transactions 		[]models.Transaction
	transactionsFlag	bool
	db           		*sql.DB
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
	if repo.usersFlag == true {
		return repo.users
	}

	result, err := repo.db.Query("SELECT * FROM users")

	if err != nil {
		log.Fatal(err)
	}

	if len(repo.users) > 0 {
		repo.users = nil
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

	repo.usersFlag = true

	return repo.users
}

func (repo *BankRepo) CreateUser(user *models.User) bool {
	_, err := repo.db.Query(
		"INSERT INTO users(login, hash_password, full_name, phone_number, email) VALUES ($1,$2,$3,$4,$5)",
				user.Login, user.HashPassword, user.FullName, user.PhoneNumber, user.Email)

	if err != nil {
		log.Fatal(err)
		return false
	}

	repo.usersFlag = false

	return true
}

func (repo *BankRepo) UpdateUser(user *models.User, id int) bool {
	_, err := repo.db.Query("UPDATE users SET login = $1, hash_password=$2, full_name=$3, phone_number=$4, email=$5 " +
		"WHERE user_id=$6", user.Login, user.HashPassword, user.FullName, user.PhoneNumber, user.Email, id)

	if err != nil {
		log.Fatal(err)
		return false
	}

	repo.usersFlag = false

	return  true
}

func (repo *BankRepo) GetUserById(id int) *models.User{
	rows, err := repo.db.Query("SELECT * FROM users WHERE user_id = $1", id)

	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Login, &user.HashPassword, &user.FullName, &user.PhoneNumber, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &user
}

func (repo *BankRepo) DeleteUserById(id int) bool {
	_, err := repo.db.Query("DELETE FROM users WHERE user_id = $1", id)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
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
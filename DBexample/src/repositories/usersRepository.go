package repositories

import (
	"GO-examples/DBexample/src/models"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type usersRepository struct {
	db *sql.DB
}

func NewUsersRepository(dbUser string, dbPassword string, dbName string) *usersRepository {
	connectionStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", connectionStr)
	panicHandler(err)
	return &usersRepository{db}
}

func (repository *usersRepository) GetAllUsers() []*models.User {
	slice := make([]*models.User, 0, 10)
	result, err := repository.db.Query("SELECT * FROM owner")
	panicHandler(err)

	for result.Next() {
		user := models.NullArgsConstructorUser()
		err := result.Scan(user.GetIndex(), user.GetAge(), user.GetName(), user.GetCitizen())
		panicHandler(err)
		slice = append(slice, user)
	}
	result.Close()
	return slice
}

func panicHandler(err error) {
	if err != nil {
		panic(err)
	}
}

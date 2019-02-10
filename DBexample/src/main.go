package main

import (
	"GO-examples/DBexample/src/repositories"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	dbUser     = "postgres"
	dbPassword = "sbprt9991"
	dbName     = "draqneel_db"
)

func main() {
	// connectionStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	// db, err := sql.Open("postgres", connectionStr)
	// panicHandler(err)

	// fmt.Println("GOOD BOOY")
	// // language=SQL
	// result, err := db.Query("SELECT * FROM owner")
	// panicHandler(err)

	// for result.Next() {
	// 	user := models.NullArgsConstructorUser()
	// 	err := result.Scan(user.GetIndex(), user.GetAge(), user.GetName(), user.GetCitizen())
	// 	panicHandler(err)
	// 	fmt.Printf("His name is %s it was %d years old, living in %s, with index %d\n", *user.GetName(), *user.GetAge(), *user.GetCitizen(), *user.GetIndex())
	// }
	// result.Close()
	// defer db.Close()
	repository := repositories.NewUsersRepository(dbUser, dbPassword, dbName)
	for _, value := range repository.GetAllUsers() {
		fmt.Println(value)
	}
}

// func panicHandler(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

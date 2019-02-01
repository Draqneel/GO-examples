package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	dbUser     = "your_user"
	dbPassword = "your_password"
	dbName     = "your_username"
)

func main() {
	connectionStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", connectionStr)
	panicHandler(err)

	fmt.Println("GOOD BOOY")
	// language=SQL
	result, err := db.Query("SELECT * FROM owner")
	panicHandler(err)

	for result.Next() {
		var index int
		var age int
		var name string
		var citizen string
		err := result.Scan(&index, &age, &name, &citizen)
		panicHandler(err)
		fmt.Printf("His name is %s it was %d years old, living in %s, with index %d\n", name, age, citizen, index)
	}
	result.Close()
	defer db.Close()
}

func panicHandler(err error) {
	if err != nil {
		panic(err)
	}
}

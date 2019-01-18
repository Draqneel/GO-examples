package models

import (
	"math/rand"
)

type User struct {
	id      int64
	name    string
	surname string
	salary  int
}

func NewUser(name string, surmane string, salary int) *User {
	return &User{rand.Int63(), name, surmane, salary}
}

func (user *User) GetSalary() int {
	return user.salary
}

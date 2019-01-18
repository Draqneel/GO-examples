package models

import (
	"math/rand"
)

type user struct {
	id      int64
	name    string
	surname string
	salary  int
}

// NewUser - return link on created user object
func NewUser(name string, surmane string, salary int) *user {
	return &user{rand.Int63(), name, surmane, salary}
}

func (user *user) GetSalary() int {
	return user.salary
}

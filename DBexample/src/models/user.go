package models

// User - model of user
type User struct {
	index   int
	age     int
	name    string
	citizen string
}

// FullArgsConstructorUser - constructor of user
func FullArgsConstructorUser(index int, age int, name string, citizen string) *User {
	return &User{index, age, name, citizen}
}

// NullArgsConstructorUser - constructor of user
func NullArgsConstructorUser() *User {
	return &User{0, 0, "", ""}
}

func (user *User) GetIndex() *int {
	return &user.index
}

func (user *User) GetAge() *int {
	return &user.age
}

func (user *User) GetName() *string {
	return &user.name
}

func (user *User) GetCitizen() *string {
	return &user.citizen
}

package models

type user struct {
	index   int
	age     int
	name    string
	citizen string
}

// FullArgsConstructorUser - constructor of user
func FullArgsConstructorUser(index int, age int, name string, citizen string) *user {
	return &user{index, age, name, citizen}
}

// NullArgsConstructorUser - constructor of user
func NullArgsConstructorUser() *user {
	return &user{0, 0, "", ""}
}

func (user *user) GetIndex() *int {
	return &user.index
}

func (user *user) GetAge() *int {
	return &user.age
}

func (user *user) GetName() *string {
	return &user.name
}

func (user *user) GetCitizen() *string {
	return &user.citizen
}

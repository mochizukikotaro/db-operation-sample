package model

type User struct {
	Id   int
	Name string
}

func CreateUsers() []User {
	return []User{}
}

func CreateUser() User {
	return User{}
}

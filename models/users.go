package models

type Address struct {
	Country string
	State   string
	Street  string
}

type User struct {
	Id        string
	Firstname string
	LastName  string
	Email     string
	Phone     string
	UserType  string
	Address   Address
}

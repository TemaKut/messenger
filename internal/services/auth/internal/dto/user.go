package dto

type User struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Username  string
	Password  *string
}

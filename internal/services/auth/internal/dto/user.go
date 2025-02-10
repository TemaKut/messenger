package dto

// TODO required and optional fields
type User struct {
	Id        string
	Username  string
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type UnregisteredUser struct {
	Username  string
	FirstName *string
	LastName  *string
	Email     *string
	Phone     *string
	Password  *string
}

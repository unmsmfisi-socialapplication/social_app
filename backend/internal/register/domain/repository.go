package domain

type IuserDBRepository interface {
	GetUserByEmail(email string) (*User, error)
	RegisterUser( phone, email, username,password string) (*User, error)
}
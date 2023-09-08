package domain

type User struct {
	Id_user  int
	Username string
	Email    string
	Password string
	Name     string
	Lastname string
	Birthday string
}

func NewUser(id_user int, username, email, password, name, lastName, birthday string) (*User, error) {

	return &User{
		Id_user:  id_user,
		Username: username,
		Email:    email,
		Password: password,
		Name:     name,
		Lastname: lastName,
		Birthday: birthday,
	}, nil
}

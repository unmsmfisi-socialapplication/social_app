package domain

type User struct {
	Username string
	Password string
	Role     string
}

type AuthSessionRepository interface {
	CheckValidSession(username string) (bool, error)
	GetJTIByUsername(username string) (string, error)
	InvalidateSession(username string) error
}

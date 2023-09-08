package infraestructure

import (
	"database/sql"
	"errors"

	"github.com/unmsmfisi-socialapplication/social_app/internal/User/domain"
	"github.com/unmsmfisi-socialapplication/social_app/internal/database"
)

type UserDBRepository struct {
	db *database.Database //*sql.DB
}

func NewUserDBRepository(database *database.Database) *UserDBRepository {
	return &UserDBRepository{db: database}
}

func (u *UserDBRepository) GetUserById(id_user int) (*domain.User, error) {
	query := `SELECT id_user,username,email,password,name,lastname,birthday
	FROM users WHERE id_user = $1`

	row := u.db.QueryRow(query, id_user)

	var user domain.User
	err := row.Scan(&user.Id_user, &user.Username, &user.Email, &user.Password, &user.Name, &user.Lastname)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserDBRepository) UpdateUser(id_user int, username, email, password, name, lastName, birthday string) (*domain.User, error) {
	// Verificar si el usuario ya existe
	existingUser, err := u.GetUserById(id_user)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("el correo electrónico ya está en uso")
	}

	// Crear un nuevo usuario
	newUser, err := domain.NewUser(id_user, username, email, password, name, lastName, birthday) // Utilizamos el correo electrónico como identificador
	if err != nil {
		return nil, err
	}

	// Guardar el nuevo usuario en la base de datos
	query := `UPDATE users SET username = $1, email = $2, password = $3, name = $4, lastName = $5, birthday = $6 WHERE id_user = $7`
	_, err = u.db.Exec(query, newUser.Username, newUser.Email, newUser.Password, newUser.Name, newUser.Lastname, newUser.Birthday, newUser.Id_user) // Usamos el correo como username
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

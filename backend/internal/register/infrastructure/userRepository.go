package infrastructure

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/unmsmfisi-socialapplication/social_app/internal/register/domain"
	"github.com/unmsmfisi-socialapplication/social_app/pkg/database"
)

type UserDBRepository struct {
	db *database.Database
}

func NewUserDBRepository(database *database.Database) (*UserDBRepository) {
	return &UserDBRepository{db: database}
}

func (u *UserDBRepository) GetUserByEmail(email string) (*domain.User, error) {
	
	query := `SELECT  phone, email, user_name, password FROM public.soc_app_users WHERE email = $1`

	row := u.db.DB.QueryRow(query, email)
	prueba,_:= u.db.DB.Exec(query, email)
	fmt.Println(prueba.RowsAffected())
	var user domain.User
	err := row.Scan( &user.Phone, &user.Email, &user.User_name, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserDBRepository) RegisterUser( phone, email, username,password string) (*domain.User, error) {

	existingUser, err := u.GetUserByEmail(email)
	
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("el correo electrónico ya está en uso")
	}


	newUser, err := domain.NewUser(phone, email, username,password) // Utilizamos el correo electrónico como identificador
	if err != nil {
		return nil, err
	}
	

	query := `INSERT INTO soc_app_users ( insertion_date,phone, email, user_name, password) VALUES (NOW(),$1, $2, $3, $4)`

	fmt.Println("Insertando usuario en la base de datos...")
	tx,er:=u.db.DB.Begin()
	if er!=nil{
		fmt.Println("Error al iniciar la transaccion")

	}
	_,err=tx.Exec(query,newUser.Phone,newUser.Email,newUser.User_name,newUser.Password) // Usamos el correo como username
	if err != nil {
		return nil, err
	}

	tx.Commit()

	
	return newUser, nil
}
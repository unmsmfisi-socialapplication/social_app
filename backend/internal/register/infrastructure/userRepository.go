package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/unmsmfisi-socialapplication/social_app/internal/register/domain"
)

type UserDBRepository struct {
	db *sql.DB
}

func NewUserDBRepository(database *sql.DB) (*UserDBRepository) {
	return &UserDBRepository{db: database}
}

func (u *UserDBRepository) GetUserByEmail(email string) (*domain.User, error) {
	
	query := `SELECT  phone, email, user_name, password FROM public.soc_app_users WHERE email = $1`

	row := u.db.QueryRow(query, email)
	prueba,_:= u.db.Exec(query, email)
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

func (u *UserDBRepository) InsertUser(newUser *domain.User) (*domain.User, error) {

	query := `INSERT INTO soc_app_users ( insertion_date,phone, email, user_name, password) VALUES (NOW(),$1, $2, $3, $4)`

	fmt.Println("Insertando usuario en la base de datos...")
	tx,er:=u.db.Begin()
	if er!=nil{
		fmt.Println("Error while starting the transaction")
	}
	
	_,err:=tx.Exec(query,newUser.Phone,newUser.Email,newUser.User_name,newUser.Password) 
	if err != nil {
		return nil, err
	}

	tx.Commit()
	
	return newUser, nil
}
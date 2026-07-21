package services

import (
	"database/sql"
	"errors"
	"gym-shop/repositories"
)

func Register(
	db *sql.DB,
	name string,
	email string,
	password string,
) error {

	user := repositories.FindUserByEmail(db, email)

	if user != nil {
		return errors.New("email already exists")
	}
	
	return repositories.CreateUser(
		db,
		name,
		email,
		password,
	)

}

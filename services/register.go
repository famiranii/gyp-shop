package services

import (
	"database/sql"
	"errors"
	"gym-shop/repositories"
	"strings"
)

func Register(
	db *sql.DB,
	name string,
	email string,
	password string,
) error {

	user := repositories.FindUserByEmail(db, email)
	
	if !strings.Contains(email, "@") {
	return errors.New("email must contain @")
}
	
	if len(password) < 8 {
		return errors.New("password is too short")
	}
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

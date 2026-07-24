package services

import (
	"database/sql"
	"errors"
	"gym-shop/models"
	"gym-shop/repositories"
)

func LoginCheck(
	db *sql.DB,
	email string,
	password string,
) (*models.User, error) {

	user := repositories.FindUserByEmail(db, email)

	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	if user.Password != password {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
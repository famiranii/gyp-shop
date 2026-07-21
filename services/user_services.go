package services

import (
	"database/sql"
	"gym-shop/models"
	"gym-shop/repositories"
)

func GetUsers(db *sql.DB) ([]models.User, error) {

	users, err := repositories.GetUsers(db)

	if err != nil {
		return nil, err
	}

	return users, nil
}
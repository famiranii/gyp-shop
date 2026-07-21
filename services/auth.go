package services

import (
	"database/sql"
	"gym-shop/repositories"
)

func LoginCheck(
	db *sql.DB,
	email string,
	password string,
) bool {
	user := repositories.FindUserByEmail(db, email)
	if user == nil {
		return false
	}
	return user.Password == password
}

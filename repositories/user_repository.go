package repositories

import (
	"database/sql"
	"gym-shop/models"
)

func FindUserByEmail(db *sql.DB, email string) *models.User {

	row := db.QueryRow(
		"SELECT id, name, email, password FROM users WHERE email=$1",
		email,
	)

	var user models.User

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil
	}

	return &user
}

func CreateUser(
	db *sql.DB,
	name string,
	email string,
	password string,
) error {

	_, err := db.Exec(
		"INSERT INTO users(name, email, password) VALUES($1, $2, $3)",
		name,
		email,
		password,
	)

	if err != nil {
		return err
	}

	return nil
}

func GetUsers(db *sql.DB) ([]models.User, error) {

	rows, err := db.Query(
		"SELECT id, name, email, password FROM users",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	// چک کردن خطاهای حین خواندن rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
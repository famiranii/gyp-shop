package services

import (
	"database/sql"
	"gym-shop/models"
	"gym-shop/repositories"
)

func GetProducts(db *sql.DB) ([]models.Product, error) {

	products, err := repositories.GetProcuts(db)

	if err != nil {
		return nil, err
	}

	return products, nil
}
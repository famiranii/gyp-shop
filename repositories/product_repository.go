package repositories

import (
	"database/sql"
	"gym-shop/models"
)


func GetProcuts(db *sql.DB) ([]models.Product, error) {

	rows, err := db.Query(
		"SELECT id, name, description, price, stock, image_url FROM products",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []models.Product

	for rows.Next() {

		var product models.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.ImageURL,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	// چک کردن خطاهای حین خواندن rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
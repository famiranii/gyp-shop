package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {

	db, err := sql.Open(
		"postgres",
		"host=localhost port=5432 user=postgres password=4080 dbname=gym-shop sslmode=disable",
	)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected")

	return db
}

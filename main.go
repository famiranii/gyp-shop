package main

import (
	"fmt"
	"gym-shop/config"
	"gym-shop/database"
	"gym-shop/handlers"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()
	db := database.Connect()
	handler := handlers.Handler{
		DB: db,
	}
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/register", handler.RegisterHandler)
	http.HandleFunc("/users", handler.UsersHandler)
	http.HandleFunc("/products", handler.ProductsHandler)

	fmt.Println("server running on port", cfg.Port)

	http.ListenAndServe(":"+cfg.Port, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Home:", r.Method, r.URL.Path)

	fmt.Fprintln(w, "Home Page")
}

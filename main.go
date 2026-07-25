package main

import (
	"fmt"
	"gym-shop/config"
	"gym-shop/database"
	"gym-shop/handlers"
	"gym-shop/middleware"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	cfg := config.LoadConfig()
	db := database.Connect()
	handler := handlers.Handler{
		DB: db,
	}
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/login", handler.LoginHandler)
	mux.HandleFunc("/register", handler.RegisterHandler)

    mux.HandleFunc("/users", middleware.Auth(handler.UsersHandler))
    mux.HandleFunc("/products", middleware.Auth(handler.ProductsHandler))

	fmt.Println("server running on port", cfg.Port)

	http.ListenAndServe(":"+cfg.Port, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home:", r.Method, r.URL.Path)

	fmt.Fprintln(w, "Home Page")
}

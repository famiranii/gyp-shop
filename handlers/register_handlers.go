package handlers

import (
	"encoding/json"
	"fmt"
	"gym-shop/services"
	"net/http"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h Handler) RegisterHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", 405)
		return
	}

	var user RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&user)
	if user.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	if user.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	if user.Password == "" {
		http.Error(w, "Password is required", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}
	err = services.Register(
		h.DB,
		user.Name,
		user.Email,
		user.Password,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "Register success")
}

package handlers

import (
	"database/sql"
	"encoding/json"
	"gym-shop/services"
	"gym-shop/utils"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Handler struct {
	DB *sql.DB
}

func (h Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", 405)
		return
	}

	var user LoginRequest

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	userInDataBase, err := services.LoginCheck(
		h.DB,
		user.Email,
		user.Password,
	)
	if err != nil {
		http.Error(w, "Wrong email or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(userInDataBase.ID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})

}

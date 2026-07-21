package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gym-shop/services"
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

	if services.LoginCheck(
		h.DB,
		user.Email,
		user.Password,
	) {
		fmt.Fprintln(w, "Login success")
		return

	}

	http.Error(w, "Wrong email or password", 401)

}

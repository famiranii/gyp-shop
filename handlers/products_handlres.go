package handlers

import (
	"encoding/json"
	"net/http"
	"gym-shop/services"
)

func (h Handler) ProductsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", 405)
		return
	}

	products, err := services.GetProducts(h.DB)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(products)
}
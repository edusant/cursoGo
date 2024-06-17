package controllers

import (
	"encoding/json"
	"net/http"

	"../models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	users := models.GetUsers() // Assuming GetUsers() retrieves users from a data source

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

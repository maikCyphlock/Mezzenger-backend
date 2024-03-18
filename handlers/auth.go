package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"db/db"
	// Use the module name instead of a relative path
)

type User struct {
	ID       int
	Username string
	Password string
}

func Register(w http.ResponseWriter, r *http.Request) {
	var u db.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	user, err := db.CreateUser(&u)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created", "id": strconv.Itoa(user.ID)})
}
func Login(w http.ResponseWriter, r *http.Request) {
	var credentials map[string]string
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username := credentials["username"]
	password := credentials["password"]

	user, err := db.GetUserByCredentials(username, password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Bienvenido"})
}

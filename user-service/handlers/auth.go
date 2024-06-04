package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-app/user-service/db"
	"social-app/user-service/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	Storage *db.Storage
}

func NewHandler(storage *db.Storage) *Handler {
	return &Handler{Storage: storage}
}

var users = make(map[string]models.User)

func (h *Handler)RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.Createdat = time.Now()
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.Password = string(hashed)
	users[user.Email] = user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func (h *Handler)LoginUser(w http.ResponseWriter, r *http.Request) {
	type login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var userData login
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := users[userData.Email]

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

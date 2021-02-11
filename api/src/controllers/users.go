package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// GetUser -
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Users."))
}

// GetUserByID -
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User by ID."))
}

// CreateUser -
func CreateUser(w http.ResponseWriter, r *http.Request) {
	req, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	if err = json.Unmarshal(req, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Init()

	if err != nil {
		log.Fatal(err)
	}

	userRepository := repositories.NewRepositoryUsers(db)
	userRepository.Create(user)
}

// UpdateUser -
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User."))
}

// DeleteUser -
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User."))
}

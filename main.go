package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = []User{
	User{Id: 1, Name: "gato", Email: "gato@soneca.com", Password: "123"},
	User{Id: 2, Name: "gato2", Email: "gato2@soneca.com", Password: "123sada"},
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/users", getUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

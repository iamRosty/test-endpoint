package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	port                string = "8080"
	minPasswordLen      int    = 8
	maxPasswordLen      int    = 256
	minNameLen          int    = 2
	usersResourcePrefix string = "/users"
)

var db []User

type User struct {
	Id       int
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	IsAdmin  bool   `json:"isadmin"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Message struct {
	Message string `json:"message"`
}

func initHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	log.Println("Creating new user...")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Message{Message: "Provided json file is invalid"})
		return
	}
	user.Id = len(db) + 1
	db = append(db, user)
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(user)
}
func main() {
	log.Printf("Starting server at port: %s\n", port)
	router := mux.NewRouter()
	router.HandleFunc(usersResourcePrefix, RegisterUser).Methods("POST")
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}

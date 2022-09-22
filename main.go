package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/mail"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	port                string = "8080"
	minPasswordLen      int    = 8
	maxPasswordLen      int    = 256
	minNameLen          int    = 2
	usersResourcePrefix string = "/users"
	connStr             string = "user=app password=pass dbname=db sslmode=disable"
)

var db sql.DB

type User struct {
	Id       int    `json:"id"`
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

func (user *User) validationUserData() string {
	var msg string
	if len(user.Name) < minNameLen {
		msg = "The minimum length of the name is at least 2 characters"
	}
	if len(user.Password) < minPasswordLen || len(user.Password) > maxPasswordLen {
		msg = "The password must be at least 8 characters long and no longer than 256 characters"
	}
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		msg = "You entered the wrong email, it should be username@hostname"
	}
	return msg
}
func Create(db *sql.DB, user *User) {
	_, err := db.Exec("insert into Users (name, surname, isadmin, email, password) values ($1, $2, $3, $4, $5)",
		user.Name, user.Surname, user.IsAdmin, user.Email, user.Password)
	if err != nil {
		log.Fatal(err)
	}
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	log.Println("Trying to create a new user...")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Message{Message: "Provided json file is invalid"})
		return
	}
	msg := user.validationUserData()
	if msg != "" {
		msgInfo := Message{Message: msg}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msgInfo)
		return
	}
	Create(&db, &user)
	log.Println("A new user was created")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(user)
}
func main() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Printf("Starting server at port: %s\n", port)
	router := mux.NewRouter()
	router.HandleFunc(usersResourcePrefix, RegisterUser).Methods("POST")
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}

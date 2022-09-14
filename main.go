package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port string = ":8080"

var db []User

type User struct {
	Id       int
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	IsAdmin  bool   `json:"isadmin"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	log.Printf("Starting server at port: %s\n", port)
	router := mux.NewRouter()
	router.HandleFunc("users", RegisterUser).Methods("POST")
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}

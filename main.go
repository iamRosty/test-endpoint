package main

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

}

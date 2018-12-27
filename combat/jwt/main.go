package main

import (
	"log"
	"net/http"
)

const (
	SecretKey = "Welcome to drzhang"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type UserCredentials struct {
	Username string	`json:"username"`
	Password string	`json:"password"`
}

type User struct {
	ID int			`json:"id"`
	Name string		`json:"name"`
	Username string	`json:"username"`
	Password string	`json:"password"`
}

type Response struct {
	Data string		`json:"data"`
}

type Token struct {
	Token string	`json:"token"`
}

func StartServer() {
	http.HandleFunc("/login",LoginHandler)
}

func main() {
	
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{"Gained access to protected resource"}
	JsonResponse(response,w)
}

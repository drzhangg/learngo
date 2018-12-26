package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Person struct {
	ID string	`json:"id,omitempty"`
	Firstname string	`json:"firstname,omitempty"`
	Lastname string		`json:"lastname,omitempty"`
	Address *Address	`json:"address,omitempty"`
}

type Address struct {
	City string		`json:"city,omitempty"`
	Province string	`json:"province,omitempty"`
}

var people []Person

func GetPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(people)
}

func GetPeople(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func PostPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people,person)
	json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, req *http.Request) {

}

func main() {
	
}

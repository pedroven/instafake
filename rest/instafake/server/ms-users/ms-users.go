package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type allUsers []User

var Users = allUsers{
	{
		ID: 1, Name: "John", Login: "john1", Password: "johnps",
	},
	{
		ID: 2, Name: "Mary", Login: "mary2", Password: "maryps",
	},
	{
		ID: 3, Name: "Lia", Login: "lia3", Password: "liaps",
	},
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Users)
}

func getOneUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	var id int
	id, _ = strconv.Atoi(key)
	user := make([]User, 0)
	var kf int
	for key, user := range Users {
		if user.ID == id {
			kf = key
			break
		}
	}
	user = append(user, Users[kf])
	json.NewEncoder(w).Encode(user)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", getUsers)
	myRouter.HandleFunc("/users/{id}", getOneUser)
	log.Fatal(http.ListenAndServe(":7777", myRouter))
}

func main() {
	handleRequests()
}

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
	NickName    string `json:"nick"`
	Password string `json:"password"`
	Description string `json:"description"`
}

type allUsers []User

var Users = allUsers{
	{
		ID: 1, Name: "John Wick", NickName: "johnwick", Password: "johnw", Description: "Don´t mess with my dog",
	},
	{
		ID: 2, Name: "Ray", NickName: "raySue", Password: "raysue", Description: "I owned the jedi",
	},
	{
		ID: 3, Name: "Kilo Ren", NickName: "kilo_ren", Password: "kiloRen", Description: "I love my Grandfather",
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

package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getUsersList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(getUsersListRepo())
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	json.NewEncoder(w).Encode(getUserRepo(id))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	deleteUserRepo(id)
	json.NewEncoder(w).Encode("Ok!")
}

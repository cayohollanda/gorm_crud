package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	loadDb()
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsersList).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

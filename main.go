package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	loadDb()
	/*user := User{
		Name:     "Cayo",
		Username: "cayo.andrade",
		Email:    "cayohollanda@hotmail.com",
		Role:     "ADMIN",
	}*/
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsersList).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

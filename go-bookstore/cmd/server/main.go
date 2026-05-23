package main

import (
	"bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Welcome to the go - bookstore")
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}

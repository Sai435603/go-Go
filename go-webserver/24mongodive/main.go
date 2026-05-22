package main

import (
	"fmt"
	"log"
	"mongo-go/routers"
	"net/http"
)

func main() {
	fmt.Println("MongoDB api")
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":8000", routers.Router()))

}

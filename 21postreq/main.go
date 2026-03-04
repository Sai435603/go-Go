package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	const myUrl = "http://localhost:3000/data"
	PerformGetRequest(myUrl)
}

func PerformGetRequest(myUrl string) {
	responseBody := strings.NewReader(`{
	    "course" : "go-lang",
		"cst" : 100000,
		"cnt" : 20,
		"diff" : "hard"
	  }`)

	response, err := http.Post(myUrl, "application/json", responseBody)
	fmt.Println(responseBody)
	if err != nil {
		panic(err)
	}

	data, _ := io.ReadAll(response.Body)

	fmt.Println(string(data))
}




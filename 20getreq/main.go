package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	const myUrl = "http://localhost:3000"
	PerformGetRequest(myUrl)
}

func PerformGetRequest(myUrl string) {
	response, err := http.Get(myUrl)

	CheckNilError(err)

	defer response.Body.Close()

	fmt.Println(response.Status, response.ContentLength)

	// way-1
	content, _ := io.ReadAll(response.Body) 

	fmt.Println(string(content))

	//way-2
	var resStr strings.Builder

	lenght, _ := resStr.Write(content)
	fmt.Println(lenght)
	fmt.Println(resStr.String())
}

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}

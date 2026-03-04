package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	const myUrl = "http://localhost:3000/postform"
	PostFormData(myUrl)
}

func PostFormData(myUrl string) {
	data := url.Values{}
	data.Add("name", "sai")
	data.Add("email", "google@gmail.com")
	data.Add("age", "20")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://codeforces.com"

func main() {
	response, err := http.Get(url)

	ChkNilErr(err)

	defer response.Body.Close()

	dataBytes, err := io.ReadAll(response.Body)
	ChkNilErr(err)
	fmt.Println(string(dataBytes))

}

func ChkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

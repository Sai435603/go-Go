package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://loc.dev:3000/learn?coursename=reactjs&paymentid=w34fkhe2983"

func main() {
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Path, result.Scheme, result.User, result.RawQuery, result.Port(), result.Hostname(), result.String())
	params := result.Query()
	fmt.Printf("%T\n%v\n", params, params)
	fmt.Println(params["coursename"], params["paymentid"][0])
	fmt.Printf("%T\n", params["paymentid"])

	urlChunks := &url.URL{
		Scheme:  "https",
		Host:    "sai.com",
		Path:    "/courses",
		RawPath: "user=sai",
	}

	anotherUrl := urlChunks.String()
	fmt.Println(anotherUrl)
}

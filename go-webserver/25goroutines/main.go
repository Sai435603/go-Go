package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mtx sync.Mutex

func main() {
	endpoints := []string{
		"https://codeforces.com",
		"https://go.dev",
		"https://github.com",
		"https://fb.com",
		"https://leetcode.com/",
	}

	for _, endpoint := range endpoints {
		go getStatusCode(endpoint)
		wg.Add(1)
	}

	wg.Wait()
}

// func greeter(str string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(str)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mtx.Lock()
		fmt.Printf("%d status code for %v\n", res.StatusCode, endpoint)
		mtx.Unlock()
	}
}

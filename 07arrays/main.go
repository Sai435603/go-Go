package main

import "fmt"

func main() {
	var arr [4]string
	arr[0] = "hello"
	arr[3] = "world"
	fmt.Println(arr)

	var arr1 = [3]string{"hell", "yea", "go"}
	fmt.Println(arr1, len(arr1), len(arr))
}

package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("world1")
	defer fmt.Println("world2")

	fmt.Println("Hello")
	MyFun()
}

func MyFun() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

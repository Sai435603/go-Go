package main

import "fmt"

func main() {
	newUser := User{"Sai", "322106410000", 21}
	fmt.Println(newUser)
	fmt.Printf("more detailed %+v\n", newUser)
	fmt.Println("printing name only", newUser.Name)
}

type User struct {
	Name string
	Id   string
	Age  int
}

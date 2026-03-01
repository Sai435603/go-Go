package main

import "fmt"

func main() {
	usr := User{"sai", "123", 20, "email@gmail.com"}
	fmt.Println(usr)
    usr.setName()
	fmt.Println(usr.name)
}

type User struct {
	name  string
	Id    string
	Age   int
	email string
}

func (u User) setName() {
	u.name = "saii"
	fmt.Println(u.name, " ", u.Age)
}

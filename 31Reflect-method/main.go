package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func main() {
	//reflect is basically use to know the type of unknown data at runtime like we can know the thype if we use [any or interface{}]
	name := "Sai"

	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.ValueOf(name))
	u := User{
		Name: "Sai",
		Age:  22,
	}
	t := reflect.TypeOf(u)
	fmt.Println(t.Kind())
	fmt.Println(t.Name())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		fmt.Println(field.Name)
	}

	//t is of type User it consist of
	// u
	// -> Type name (User)
	// -> Number of Fields (2)
	// ->Field 0 (Name)
	// ->Field 1 (Age)
	// like wise
}

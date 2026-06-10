package main

import (
	"errors"
	"fmt"
)

func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("Divide by zero error")
	}
	return float64(a) / float64(b), nil
}
func main() {
	//error wrapping mean adding extra context to the error using fmt.Errorf()
	_, err := Divide(69, 0)
	if err != nil {
		//%w to wrap the error inside a new error jz like a %v
		newError := fmt.Errorf("Got the error from Divide function : %w", err)
		// fmt.Println(newError)
		panic(newError)
	}
	//recover() can only be used in the defer functions , it is like try catch in other languages, panic throws an err recover catches it
	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println(msg)
		}
	}()
}

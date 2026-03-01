package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("study abt switch and case")
	random := rand.IntN(10) + 1
	fmt.Println(random)

	switch random {
	case 1:
		fmt.Println("hell0")
	case 2:
		fmt.Println("hell0")
	case 3:
		fmt.Println("hell0")
	default:
		fmt.Println("default")
	}
}

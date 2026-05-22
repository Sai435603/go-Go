package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	greet()
	fmt.Println(greeting())
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	x, _ := strconv.Atoi(strings.TrimSpace(input))
	input1, _ := reader.ReadString('\n')
	y, _ := strconv.Atoi(strings.TrimSpace(input1))
	fmt.Println(addTwoNumberForMe(x, y))
	fmt.Println(addPro(1, 2, 3, 4, 5))
}

func addTwoNumberForMe(x int, y int) int {
	return x + y
}

func addPro(x ...int) int {
	var sum int = 0
	for _, ele := range x {
		sum += ele
	}
	return sum
}
func greet() {
	fmt.Println("Hey sai all good :) ?? ")
}

func greeting() string {
	return "Nope I'm not good"
}

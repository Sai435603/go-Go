package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to my app")
	fmt.Println("Rate this code from 1 to 5")

	// reading dt from user
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating ", input)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to the entered rating ", numRating+1)
	}

}

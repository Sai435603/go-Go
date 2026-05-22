package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the your age: ")
	input, _ := reader.ReadString('\n')
	age, _ := strconv.Atoi(strings.TrimSpace(input))

    if(age < 18) {
		 fmt.Println("Under Age")
	} else if(age > 18) {
		 fmt.Println("Over Age")
	} else { 
		fmt.Println("Perfect Age")
	}
}


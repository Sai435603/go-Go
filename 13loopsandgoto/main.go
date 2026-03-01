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
	arr := []string{}
	fmt.Println("Enter the size")
	input, _ := reader.ReadString('\n')
	sz, _ := strconv.Atoi(strings.TrimSpace(input))
	// fmt.Println(sz);
	fmt.Printf("Enter Array elements %v\n", sz)

	for i := 0; i < sz; i++ {
		x, _ := reader.ReadString('\n')
		arr = append(arr, strings.TrimSpace(x))
	}

	fmt.Println(arr)

	// loops
	//typ 1
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v ", arr[i])
	}
	fmt.Println()

	//typ 2

	for ind := range arr {
		fmt.Printf(" %v", arr[ind])
	}

	fmt.Println()

	//typ3

	for ind, ele := range arr {
		fmt.Printf("[%v, %v] ", ind, ele)
	}
	fmt.Println()

	//typ 4 goto included

	var ind int = 3
	for ind >= 0 {
		fmt.Println(arr[ind])
		if ind == 1 {
			goto lable
		}
		ind--
	}

lable:
	fmt.Println("from lable")

}

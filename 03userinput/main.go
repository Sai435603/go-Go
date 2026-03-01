package main

import (
	"bufio"
	"fmt"
	"os"
)

func main (){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the input string")
	input,_ := reader.ReadString('\n');
	fmt.Println("Enter data is: ", input);
}
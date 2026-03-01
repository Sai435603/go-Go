package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	content := "Hey how are you :)"
	file, err := os.Create("./testfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println(length)

	// Move pointer to beginning
	file.Seek(0, 0)

	reader := bufio.NewReader(file)
	inp, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Println(strings.TrimSpace(inp))

	//tp2 with ioutil
	dt, err := ioutil.ReadFile("testfile.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dt))

}

package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	for true {
		clear()
		h, m, s := time.Now().Clock()
		fmt.Printf("%v : %v : %v \n", h, m, s)
		time.Sleep(1 * time.Second)
	}
}

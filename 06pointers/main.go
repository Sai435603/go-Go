package main

import "fmt"

func main(){
	 var ptr *int
	 fmt.Println("initial ptr address is ", ptr);
	//  fmt.Println("ptr initial val is ", *ptr)
     
	 val := 10
	 ptr1 := &val
	 fmt.Println("add of val var is ", &val)
     fmt.Println("add of ptr1 is ", ptr1)
	 fmt.Println("actual val of ptr1 is ", *ptr1);
}
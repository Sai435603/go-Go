package main

import "fmt"
// myvar1 := "myvarDatatype"  this will give error have to mention the var and type

// Node:- if we use to declare the a var with first letter capital then it is treated as the PUBLIC var

var MyPublicVar int = 10000000007

func main(){
	 // string type
	 var name string = "Sai"
	 fmt.Println(name);
	 fmt.Printf("This variable is of type %T \n", name);
    
	 // integer type 
	 var num1 int = 100
	 var num2 uint8 = 200
	 var num3 uint32 = 400
	 var num4 uint64 = 500
	 var num5 uint16 = 600
	 fmt.Println(num1, num2, num3, num4, num5);
	 fmt.Printf("This variable is of type %T %T %T %T %T\n", num1, num2, num3, num4, num5);

	//  float type
	var fnum1 float32 = 10.082340735734234
	var fnum2 float64 = 10.082340735734234
	fmt.Println(fnum1, fnum2);
	fmt.Printf("This variables is of types %T %T \n", fnum1, fnum2);
     
	// different styles of declaring a var
	// 1. without type mentioning
	var myvar = "anyType"
	fmt.Println(myvar)
	//myvar = 10 will get error

	// 2. without var and data-type need to use := symbol and can declare this kind or this style in global
	myvar1 := "myvarDatatype"
	fmt.Println(myvar1)
    
	// public var  
	fmt.Printf("This is a public Var %d \n", MyPublicVar)

}
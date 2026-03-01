package main

import "fmt"

func main() {
	 fmt.Println("maps in go lang")

	 var months = make(map[string]string)
     months["jan"] = "january"
	 months["feb"] = "february"
	 months["mar"] = "march"
	 months["apr"] = "april"

	 for key,val := range months{
		 fmt.Println(key," ",val)
	 }
     
	 fmt.Println(months)

	 delete(months,"feb")

	 fmt.Println(months)
}
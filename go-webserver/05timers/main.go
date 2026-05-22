package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("a bit about the fr timers ")
	//to get the current time
	fmt.Println("The current time is ", time.Now())
	// we can format the date by using the method called Format that is from time pkg
	// rules date : "01-02-2006"  day : "Monday" time : "15:04:05"
	currentTime := time.Now()
	fmt.Println(currentTime.Format("01-02-2006 Monday 15:04:05"))

	createdTime := time.Date(2020, time.August, 10, 12, 32, 32, 20, time.UTC)
	fmt.Println("created time is: ", createdTime)
	fmt.Println("Formatted created time is ", createdTime.Format("01-02-2006 Monday 15:04:05"))
}

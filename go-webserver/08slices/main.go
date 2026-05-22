package main

import (
	"fmt"
	"sort"
)

func main() {
	// this is how we declare a slice
	var slc = []string{}
	fmt.Printf("type of slc is %T \n", slc)
	// 232574

	var slc1 = []string{"hell", "apple"}
	slc1 = append(slc1, "mango", "banana")
	fmt.Println(slc1)

	var slc2 = append(slc1[:2])
	fmt.Println(slc2)

	var nums = make([]int, 5)
	nums[0] = 2
	nums[1] = 1
	nums[2] = 4
	nums[3] = 3
	nums[4] = -20
	fmt.Println(sort.IntsAreSorted(nums))
	sort.Ints(nums)
	fmt.Println(sort.IntsAreSorted(nums))

	// removind 2 number from nums using append
	nums = append(nums[:2], nums[3:]...)
	fmt.Println(nums)

}

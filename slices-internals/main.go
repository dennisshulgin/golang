package main

import (
	"fmt"
)

func main() {
	nums := make([]int, 0, 5)

	for i := 0; i < 11; i++ {
		nums = append(nums, i)
		fmt.Printf("len=%v, cap=%v, nums=%v \n", len(nums), cap(nums), nums)
	}

	a := []int{1, 2, 3, 4, 5}
	b := a[1:2]
	fmt.Println(len(b) == 1)
	fmt.Println(b[0] == 2)

	c := a[:2]
	c = append(c, 100)
	fmt.Println(a)
	fmt.Println(c)
	c = append(c, 100)
	c = append(c, 100)
	c = append(c, 100)
	c = append(c, 100)
	fmt.Println(a)
	fmt.Println(c)
}

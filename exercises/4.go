package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f_nums := []int{0}
	
	return func() int {
		length := len(f_nums)
		if length == 1 {
			f_nums = append(f_nums, 1)
		} else {
			f_nums = append(f_nums, f_nums[length-2] + f_nums[length-1])
		}
		return f_nums[length-1]
	}
}

func four() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

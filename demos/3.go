package main

import "fmt"

func three() {
	defer fmt.Println("done")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

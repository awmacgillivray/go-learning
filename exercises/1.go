package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := range 10 {
		diff := (z*z - x) / (2*z)
		if math.Abs(diff) < 1e-15 {
			fmt.Println(i, "iterations")
			return z
		}
		z -= diff
		// fmt.Println(z)
	}
	return z
}

func one() {
	fmt.Println(Sqrt(1))
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(3))
	fmt.Println()
	fmt.Println(math.Sqrt(1))
	fmt.Println(math.Sqrt(2))
	fmt.Println(math.Sqrt(3))
}

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt2(x float64) (float64, error) {
	if x < 0.0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	for i := range 10 {
		diff := (z*z - x) / (2*z)
		if math.Abs(diff) < 1e-15 {
			fmt.Println(i, "iterations")
			return z, nil
		}
		z -= diff
		// fmt.Println(z)
	}
	return z, nil
}

func six() {
	fmt.Println(Sqrt2(2))
	fmt.Println(Sqrt2(-2))
}

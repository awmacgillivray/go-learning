package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func seven() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex3{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex3 struct {
	X, Y float64
}

func (v *Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

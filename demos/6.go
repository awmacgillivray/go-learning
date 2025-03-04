package main

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex2) Scale(f float64) {  // need pointer in order to modify receiver itself, return vs change values
	v.X = v.X * f
	v.Y = v.Y * f
}

func six() {
	v := Vertex2{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))
	
	v.Scale(10)
	fmt.Println(v.Abs())
}

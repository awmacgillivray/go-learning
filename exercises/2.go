package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for i := range a {
		a[i] = make([]uint8, dx)
	}

	for y := range dy {
		for x := range dx {
			intValue := x^y
			a[y][x] = uint8(intValue)
		}
	}
	return a
}

func two() {
	pic.Show(Pic)
}

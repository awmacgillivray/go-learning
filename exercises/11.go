package main

import (
	"fmt"
	"slices"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	ch <- t.Value
	Walk(t.Left, ch)
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	vals1 := []int{}
	vals2 := []int{}
	
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		vals1 = append(vals1, <-ch1)
		vals2 = append(vals2, <-ch2)
	}

	slices.Sort(vals1)
	slices.Sort(vals2)
	return slices.Equal(vals1, vals2)
}

func eleven() {
	s := Same(tree.New(1), tree.New(2))
	fmt.Println("Same?", s)
}

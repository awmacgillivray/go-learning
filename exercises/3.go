package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	word_count := make(map[string]int)
	words := strings.Split(s, " ")

	for _, v := range words {
		word_count[v] += 1
	}
	
	return word_count
}

func three() {
	wc.Test(WordCount)
}

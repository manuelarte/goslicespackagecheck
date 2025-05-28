package main

import (
	"fmt"
	"slices"
)

func main() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Printf("Max(a)=%v\n", maxValue(a))
	fmt.Printf("slices.Max(a)=%v\n", slices.Max(a))
}

func maxValue(a []int) int {
	var m int
	for _, value := range a {
		if m < value {
			m = value
		}
	}
	return m
}

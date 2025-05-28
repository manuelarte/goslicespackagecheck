package main

import (
	"fmt"
	"slices"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9, 10}

	fmt.Printf("concat(a, b)=%v\n", concat(a, b))
	fmt.Printf("slices.Concat(a, b)=%v\n", slices.Concat(a, b))
}

func concat(a, b []int) []int {
	var c []int
	for _, value := range a {
		c = append(c, value)
	}
	for _, value := range b {
		c = append(c, value)
	}
	return c
}

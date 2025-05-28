package main

import (
	"fmt"
	"slices"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}

	fmt.Printf("Equal(a, b)=%v\n", equal(a, b))
	fmt.Printf("slices.Equal(a, b)=%v\n", slices.Equal(a, b))
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	n := len(a)
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

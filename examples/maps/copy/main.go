package main

import (
	"fmt"
	"maps"
)

func main() {
	a := map[int]string{1: "1", 2: "2", 3: "3", 4: "4", 5: "5"}

	b := make(map[int]string)
	b2 := make(map[int]string)

	copyMap(b, a)
	fmt.Printf("copyMap(b, a) -> %v\n", b)
	maps.Copy(b2, a)
	fmt.Printf("maps.Copy(b2, a) -> %v\n", b2)
}

func copyMap(dst map[int]string, src map[int]string) {
	for key, value := range src {
		dst[key] = value
	}
}

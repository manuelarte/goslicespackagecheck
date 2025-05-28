package max

import "log"

func concatInt() []int {
	a := []int{4, 3, 2, 1}
	var b []int
	for _, value := range a { // want `this for loop can be replaced by slices.Concat`
		b = append(b, value)
	}
	return b
}

func concatAndLogInt() []int {
	a := []int{4, 3, 2, 1}
	var b []int
	for _, value := range a {
		b = append(b, value)
		log.Println(value)
	}
	return b
}

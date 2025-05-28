package max

func maxInt() int {
	a := []int{4, 3, 2, 1}
	var maxValue int
	for _, value := range a { // want `this for loop can be replaced by slices.Max`
		if value >= maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func maxFloat32() float32 {
	a := []float32{4.0, 2.0, 5.1, 0}
	var maxValue float32
	for _, value := range a { // want `this for loop can be replaced by slices.Max`
		if value >= maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func maxIntIge() int {
	a := []int{4, 3, 2, 1}
	var maxValue int
	for i := 0; i < len(a); i++ {
		if a[i] >= maxValue {
			maxValue = a[i]
		}
	}
	return maxValue
}

func maxIntIg() int {
	a := []int{4, 3, 2, 1}
	var maxValue int
	for i := 0; i < len(a); i++ {
		if a[i] > maxValue {
			maxValue = a[i]
		}
	}
	return maxValue
}

func otherFunc() int {
	a := []int{4, 3, 2, 1}
	var maxValue int
	for i := 0; i < len(a); i++ {
		if a[i] > 2 {
			maxValue = a[i]
		}
	}
	return maxValue
}

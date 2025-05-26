package max

func maxInt() int {
	a := []int{4, 3, 2, 1}
	var int maxValue
	for _, value := range a { // want `this for loop can be replaced by slices.Max`
		if value >= maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func maxFloat32() int {
	a := []float32{4.0, 2.0, 5.1, 0}
	var float32 maxValue
	for _, value := range a { // want `this for loop can be replaced by slices.Max`
		if value >= maxValue {
			maxValue = value
		}
	}
	return maxValue
}

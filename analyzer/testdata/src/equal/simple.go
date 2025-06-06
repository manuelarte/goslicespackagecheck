package main

func AreStringArrayEqualOneFieldList(a, b []string) bool { // want `the function AreStringArrayEqualOneFieldList can be replaced by slices.Equal`
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func AreStringArrayEqual(a []string, b []string) bool { // want `the function AreStringArrayEqual can be replaced by slices.Equal`
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Equals[T comparable](a, b []T) bool { // want `the function Equals can be replaced by slices.Equal`
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func EqualsReverseCheck[T comparable](a, b []T) bool { // want `the function EqualsReverseCheck can be replaced by slices.Equal`
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if b[i] != a[i] {
			return false
		}
	}
	return true
}

func EqualInsideIf[T comparable](a, b []T) bool { // want `the function EqualInsideIf can be replaced by slices.Equal`
	if len(a) == len(b) {
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
	}
	return false
}

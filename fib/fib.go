package fib

// SliceMax Returns a slice of fibinoacci numbers
func SliceMax(max int) []int {
	fibslice := []int{0, 1}
	fib := 1
	idx := 2
	for fib < max {
		fibslice = append(fibslice, fib)
		fib += fibslice[idx-1]
		idx++
	}
	return fibslice
}

// SliceMax64 Return a slice of fibinoacci numbers int64 version
func SliceMax64(max int64) []int64 {
	fibslice := []int64{0, 1}
	fib := int64(1)
	idx := 2
	for fib < max {
		fibslice = append(fibslice, fib)
		fib += fibslice[idx-1]
		idx++
	}
	return fibslice
}

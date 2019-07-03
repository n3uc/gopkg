package fib

// Return a slice of fibinoacci numbers
// NOTE: contains a 0 at element 0 so indexing is correct
func SliceMax(max int) []int {
	fibslice := []int{0, 1, 2}
	fib := 3
	idx := 3
	for fib < max {
		fib = fibslice[idx-2] + fibslice[idx-1]
		fibslice = append(fibslice, fib)
		idx++
	}
	return fibslice
}


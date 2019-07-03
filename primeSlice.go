// Package prime_dns provides a way to generate a slice of prime numbers
package primeSlice

// GetSlice will return a slice containing the prime numbers from 2 to limit
// It returns the slice
func GetSlice(limit int) []int {
	limit = limit + 1
	bools := make([]bool, limit)
	for k := 2; k*k <= limit; k++ {
		if bools[k] != true {
			for ix := k * k; ix < limit; ix += k {
				bools[ix] = true
			}
		}
	}
	primes := []int{}
	for ix := 2; ix < limit; ix++ {
		if bools[ix] != true {
			primes = append(primes, ix)
		}
	}
	return primes
}



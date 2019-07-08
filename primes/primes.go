// Package primes provides a way to generate a slice of prime numbers.
package primes

// UpTo will return a slice containing the prime numbers from 2 to limit.
// It returns the slice.
func UpTo(limit int) []int {
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

// GetSlice will return a slice containing the prime numbers from 2 to limit.
// It returns the slice.
// Deprecated: Exists for historical compatibility and should not be used. Use UpTo instead
func GetSlice(limit int) []int {
	return UpTo(limit)
}

// MaxFactor returns the largest prime factor of a number
func MaxFactor(n int) int {
	d := 2
	max := 0
	for n > 1 {
		for n%d == 0 {
			if d > max {
				max = d
			}
			n /= d
		}
		d++
		if d*d > n {
			if n > max {
				max = n
			}
			break
		}
	}
	return max
}

// Factors returns a slice containing the Prime Factors of n.
// It may contain more than one of any given prime. It lists
// all prime factors on the number.
func Factors(n int) []int {
	d := 2
	primes := []int{}
	for n > 1 {
		for n%d == 0 {
			primes = append(primes, d)
			n /= d
		}
		d++
		if d*d > n && n > 1 {
			primes = append(primes, n)
			break
		}
	}
	return primes
}

// IsPrime returns true if n is Prime
func IsPrime(n int) bool {
	d := 2
	cnt := false
	for n > 1 {
		for n%d == 0 {
			if cnt {
				return false
			}
			cnt = true
			n /= d
		}
		d++
		if d*d > n && n > 1 {
			if cnt {
				return false
			}
			cnt = true
			break
		}
	}
	return cnt
}

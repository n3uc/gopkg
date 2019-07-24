package nums

import (
	"strconv"
	"strings"
)

// Min returns the Smallest int of all params
func Min(n ...int) int {
	min := n[0]
	for _, num := range n {
		if num < min {
			min = num
		}
	}
	return min
}

// Max returns the Largest int of all params
func Max(n ...int) int {
	max := n[0]
	for _, num := range n {
		if num > max {
			max = num
		}
	}
	return max
}

// StringMul multiplies each char when converted with Atoi and returns the result
func StringMul(st string) int {
	a := strings.Split(st, "")
	m := 1
	for _, c := range a {
		i, _ := strconv.Atoi(c)
		m *= i
	}
	return m
}

// SumSlice sums up all elements in an int array
func SumSlice(list []int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum
}

// MulSlice sums up all elements in an int array
func MulSlice(list []int) int {
	sum := 1
	for _, v := range list {
		sum *= v
	}
	return sum
}

// GCD greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

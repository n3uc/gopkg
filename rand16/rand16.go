package rand16

var seed uint16 = 1

// Returns a random Value from 0 to i
func Value(i uint16) uint16 {
	seed ^= seed << 7
	seed ^= seed >> 9
	seed ^= seed << 8
	return uint16((uint32(seed) * uint32(i)) >> 16)
}

// Reseed the generator
func Seed(i uint16) {
  seed = i
}


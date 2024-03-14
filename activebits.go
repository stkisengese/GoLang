package piscine

func ActiveBits(n int) int {
	count := 0
	for n > 0 { // loop until n becomes 0
		count += n & 1 // check least significant bit and add to the count
		n >>= 1        // shift bits to right
	}
	return count
}

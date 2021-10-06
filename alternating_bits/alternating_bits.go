package alternatingbits

/*
	Given a positive integer, check whether it has alternating bits: namely,
	if two adjacent bits will always have different values.
*/
func hasAlternatingBits(n int) bool {
	result := true
	for n > 0 {
		// check if bit 0 is 1
		bit0 := n & 1
		// check if bit 1 is 1
		n1 := n >> 1
		bit1 := n1 & 1
		if bit0 == bit1 {
			result = false
			break
		}
		n = n >> 1
	}
	return result
}

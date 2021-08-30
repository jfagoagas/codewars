package countingbits

/*
How many 1-bits are in its binary representation?
Let's say n's binary representation has k significant bits indexed from 1 to k.
What are the respective positions (i.e., in ascending order) of each 1-bit?
The performance is really important in this challenge
*/
func counting_bits(number int) []int {
	onesPositions := []int{}
	position := 0
	for number > 0 {
		if is1 := number & 1; is1 == 1 {
			onesPositions = append(onesPositions, position)
		}
		number = number >> 1
		position++
	}
	result := append([]int{len(onesPositions)}, onesPositions...)
	return result
}

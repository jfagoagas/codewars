package twosum

func twoSum(numbers []int, target int) (result [2]int) {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				result[0] = i
				result[1] = j
				break
			}
		}
	}
	return
}

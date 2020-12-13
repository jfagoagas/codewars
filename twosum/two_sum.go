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

/*
func twoSum(numbers []int, target int) (result int) {
 if len(numbers) == 1 {
  return 0
 }
 for i := 1; i < len(numbers); i++ {
  if numbers[i] + numbers[0] == target {
   return numbers[i] * numbers[0]
  }
 }
 return twoSum(numbers[1:], target)
}
*/

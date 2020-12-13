package twosum

import (
	"strconv"
	"testing"
)

var validateTwoSum = []struct {
	input    []int
	target   int
	expected [2]int
}{
	{[]int{0, 1, 2}, 3, [2]int{1, 2}},
	{[]int{1234, 5678, 9012}, 14690, [2]int{1, 2}},
	{[]int{2, 2, 3}, 4, [2]int{0, 1}},
	{[]int{1721, 979, 366, 299, 675, 1456}, 2020, [2]int{0, 3}},
}

func TestTwoSum(t *testing.T) {
	for i, tt := range validateTwoSum {
		t.Run("Test"+strconv.Itoa(i), func(t *testing.T) {
			got := twoSum(tt.input, tt.target)
			if got != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, got)
			}
		})
	}
}

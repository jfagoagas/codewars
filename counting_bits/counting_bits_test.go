package countingbits

import (
	"fmt"
	"testing"
)

var validateCountingBits = []struct {
	input  int
	result []int
}{
	{161, []int{3, 0, 5, 7}},
	{1, []int{1, 0}},
}

func TestCountingBits(t *testing.T) {
	for _, tt := range validateCountingBits {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			got := counting_bits(tt.input)
			for i := range got {
				if got[i] != tt.result[i] {
					t.Errorf("Expected: %v, got: %v", tt.result, got)
				}
			}
		})
	}
}

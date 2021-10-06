package alternatingbits

import (
	"strconv"
	"testing"
)

var validateAlternatingBits = []struct {
	input    int
	expected bool
}{
	{5, true},
	{7, false},
	{11, false},
	{10, true},
	{3, false},
}

func TestAlternatingBits(t *testing.T) {
	for _, tt := range validateAlternatingBits {
		t.Run(strconv.Itoa(tt.input), func(t *testing.T) {
			got := hasAlternatingBits(tt.input)
			if got != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, got)
			}
		})
	}
}

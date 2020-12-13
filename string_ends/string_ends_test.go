package string_ends

import (
	"testing"
)

var validateStringEnds = []struct {
	input    string
	ending   string
	expected bool
}{
	{"abbc", "bc", true},
	{"abc", "d", false},
	{"", "", true},
	{" ", "", true},
	{"abc", "c", true},
	{"banana", "ana", true},
	{"a", "z", false},
	{"", "t", false},
	{"$a = $b + 1", "+1", false},
	{"    ", "   ", true},
	{"1oo", "100", false},
}

func TestStringEnds(t *testing.T) {
	for _, tt := range validateStringEnds {
		t.Run(tt.input, func(t *testing.T) {
			got := stringEnds(tt.input, tt.ending)
			if got != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, got)
			}
		})
	}
}

package validateparenthesis

import (
	"testing"
)

var validateParenthesisTests = []struct {
	input    string
	expected bool
}{
	{"()", true},
	{")", false},
	{"", true},
	{"(((((())))))", true},
}

func TestValidParenthesis(t *testing.T) {
	for _, tt := range validateParenthesisTests {
		t.Run(tt.input, func(t *testing.T) {
			got := validateParenthesis(tt.input)
			if got != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, got)
			}
		})
	}
}

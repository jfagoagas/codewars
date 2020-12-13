package splitstrings

import (
	"reflect"
	"testing"
)

var validateSplitStrings = []struct {
	input    string
	expected []string
}{
	{"abc", []string{"ab", "c_"}},
	{"dawsd", []string{"da", "ws", "d_"}},
	{"awsaws", []string{"aw", "sa", "ws"}},
}

func TestSplitStrings(t *testing.T) {
	for _, tt := range validateSplitStrings {
		t.Run(tt.input, func(t *testing.T) {
			got := splitStrings(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Expected: %v, got: %v", tt.expected, got)
			}
		})
	}
}

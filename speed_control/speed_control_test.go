package speedcontrol

import (
	"testing"
)

var validateSpeedControl = []struct {
	seconds  int
	sections []float64
	expected int
}{
	{20, []float64{0.0, 0.23, 0.46, 0.69, 0.92, 1.15, 1.38, 1.61}, 41},
	{12, []float64{0.0, 0.11, 0.22, 0.33, 0.44, 0.65, 1.08, 1.26, 1.68, 1.89, 2.1, 2.31, 2.52, 3.25}, 219},
}

func TestSplitStrings(t *testing.T) {
	for _, tt := range validateSpeedControl {
		t.Run("A", func(t *testing.T) {
			got := gps(tt.seconds, tt.sections)
			if got != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, got)
			}
		})
	}
}

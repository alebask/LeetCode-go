package medium

import (
	"fmt"
	"testing"
)

func Test_fractionToDecimal(t *testing.T) {
	tests := []struct {
		n        int
		d        int
		expected string
	}{
		{n: 1, d: 2, expected: "0.5"},
		//{n: 2, d: 1, expected: "2"},
		//{n: 4, d: 333, expected: "0.(012)"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%v / %v = %v", tt.n, tt.d, tt.expected)
		t.Run(testName, func(t *testing.T) {
			actual := fractionToDecimal(tt.n, tt.d)
			if actual != tt.expected {
				t.Errorf("expected %v, actual %v", tt.expected, actual)
			}
		})
	}
}

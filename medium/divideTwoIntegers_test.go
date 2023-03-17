package medium

import (
	"fmt"
	"testing"
)

func Test_divideTwoIntegers(t *testing.T) {
	tests := []struct {
		dividend int
		divisor  int
		expected int
	}{
		{dividend: 10, divisor: 3, expected: 3},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d / %d = %d", tt.dividend, tt.divisor, tt.expected)
		t.Run(testname, func(t *testing.T) {
			actual := divide(tt.dividend, tt.divisor)
			if actual != tt.expected {
				t.Errorf("expected %d, actual %d", tt.expected, actual)
			}
		})
	}
}

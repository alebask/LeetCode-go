package medium

import (
	"fmt"
	"testing"
)

func Test_multiply(t *testing.T) {

	tests := []struct {
		num1     string
		num2     string
		expected string
	}{
		{num1: "0", num2: "999", expected: "0"},
		{num1: "999", num2: "999", expected: "998001"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s * %s = %s", tt.num1, tt.num2, tt.expected)
		t.Run(testname, func(t *testing.T) {
			actual := multiply(tt.num1, tt.num2)
			if actual != tt.expected {
				t.Errorf("expected %s, actual %s", tt.expected, actual)
			}
		})
	}

}

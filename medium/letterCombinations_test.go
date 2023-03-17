package medium

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func Test_letterCombinations(t *testing.T) {
	tests := []struct {
		digits   string
		expected []string
	}{
		{digits: "23", expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s --> %s", tt.digits, tt.expected)
		t.Run(testname, func(t *testing.T) {
			actual := letterCombinations1(tt.digits)
			if !slices.Equal(actual, tt.expected) {
				t.Errorf("expected %s, actual %s", tt.expected, actual)
			}
		})
	}
}

package medium

import (
	"fmt"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {

	tests := []struct {
		s        string
		expected string
	}{
		{s: "babad", expected: "bab"},
		{s: "cbbd", expected: "bb"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s --> %s", tt.s, tt.expected)
		t.Run(testname, func(t *testing.T) {
			actual := longestPalindrome(tt.s)
			if actual != tt.expected {
				t.Errorf("expected %s, actual %s", tt.expected, actual)
			}
		})
	}
}

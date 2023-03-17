package medium

import (
	"fmt"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {

	tests := []struct {
		s        string
		expected int
	}{
		{s: "abcabcbb", expected: 3},
		{s: "bbbbb", expected: 1},
		{s: "pwwkew", expected: 3},
		{s: "abba", expected: 2},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s --> %d", tt.s, tt.expected)
		t.Run(testname, func(t *testing.T) {
			actual := lengthOfLongestSubstring(tt.s)
			if actual != tt.expected {
				t.Errorf("expected %d, actual %d", tt.expected, actual)
			}
		})
	}
}

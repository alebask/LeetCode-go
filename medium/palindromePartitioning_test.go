package medium

import (
	"fmt"
	"testing"
)

func Test_palindromePartitioning(t *testing.T) {
	tests := []struct {
		s        string
		expected [][]string
	}{
		{s: "aab", expected: [][]string{{"a", "a", "b"}, {"aa", "b"}}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s --> %s", tt.s, tt.expected)
		t.Run(testname, func(t *testing.T) {
			actual := palondromePartitioningChannels(tt.s)

			if len(actual) != len(tt.expected) {
				t.Errorf("expected len %v, actual len %v", len(tt.expected), len(actual))
			}
		})
	}
}

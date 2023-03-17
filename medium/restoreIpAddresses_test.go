package medium

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func Test_restorIpAddresses(t *testing.T) {
	tests := []struct {
		s        string
		expected []string
	}{
		{s: "25525511135", expected: []string{"255.255.11.135", "255.255.111.35"}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s --> %s", tt.s, tt.expected)
		t.Run(testname, func(t *testing.T) {
			actual := restoreIpAddresses(tt.s)
			if !slices.Equal(actual, tt.expected) {
				t.Errorf("expected %s, actual %s", tt.expected, actual)
			}
		})
	}
}

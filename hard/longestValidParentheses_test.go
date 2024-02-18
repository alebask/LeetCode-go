package hard

import (
	"fmt"
	"testing"
)

func Test_longestValidParentheses(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		//{s: "(()", expected: 2},
		//{s: ")()())", expected: 4},
		{s: ")()())()()()((((()))))))))))(((()))())", expected: 16},
		//{s: "", expected: 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s --> %v", tt.s, tt.expected)
		t.Run(testname, func(t *testing.T) {
			actual := longestValidParentheses(tt.s)
			if actual != tt.expected {
				t.Errorf("expected %v, actual %v", tt.expected, actual)
			}
		})
	}
}

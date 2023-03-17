package easy

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func Test_twoSum(t *testing.T) {
	var tests = []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d --> %d", tt.nums, tt.target)
		t.Run(testname, func(t *testing.T) {
			actual := twoSum(tt.nums, tt.target)
			if !slices.Equal(actual, tt.expected) {
				t.Errorf("expected %d, actual %d", actual, tt.expected)
			}
		})
	}
}

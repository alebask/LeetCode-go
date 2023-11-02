package medium

import (
	"fmt"
	"testing"
)

func Test_reorderRoutes(t *testing.T) {

	concat := func(arr ...[]int) [][]int {
		return append([][]int(nil), arr...)
	}

	tests := []struct {
		n           int
		connections [][]int
		expected    int
	}{
		{n: 6, expected: 3, connections: concat([]int{0, 1}, []int{1, 3}, []int{2, 3}, []int{4, 0}, []int{4, 5})},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v --> %v", tt.n, tt.expected)
		t.Run(testname, func(t *testing.T) {
			actual := minReorder(tt.n, tt.connections)
			if actual != tt.expected {
				t.Errorf("expected %v, actual %v", tt.expected, actual)
			}
		})
	}
}

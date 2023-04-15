package medium

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func Test_oddEvenLinkedList(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 3, 5, 2, 4}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v --> %v", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			head := oddEvenList(sliceToList(tt.input))
			actual := listToSlice(head)
			if !slices.Equal(actual, tt.expected) {
				t.Errorf("expected %d, actual %d", tt.expected, actual)
			}
		})
	}
}

func sliceToList(nums []int) *ListNode {
	ss := &ListNode{}
	s := ss
	for _, n := range nums {
		s.Next = &ListNode{Val: n}
		s = s.Next
	}
	return ss.Next
}
func listToSlice(head *ListNode) []int {
	res := make([]int, 0)
	n := head
	for n != nil {
		res = append(res, n.Val)
		n = n.Next
	}
	return res
}

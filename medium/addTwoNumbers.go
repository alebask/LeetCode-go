package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	sentinel := ListNode{Val: 0}

	digit := 0
	n := &sentinel

	for l1 != nil || l2 != nil {

		val := digit
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}

		digit = val / 10
		val = val % 10

		n.Next = &ListNode{Val: val}

		n = n.Next
	}

	if digit == 1 {
		n.Next = &ListNode{Val: 1}
	}

	return sentinel.Next
}

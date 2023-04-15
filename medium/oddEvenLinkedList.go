package medium

func oddEvenList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	odd := &ListNode{}
	even := &ListNode{}
	second := head.Next
	n := head
	i := 1
	for n != nil {
		if i%2 == 0 {
			even.Next = n
			even = even.Next
		} else {
			odd.Next = n
			odd = odd.Next
		}
		n = n.Next
		i++
	}

	even.Next = nil
	odd.Next = second

	return head
}

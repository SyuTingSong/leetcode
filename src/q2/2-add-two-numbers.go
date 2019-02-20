package q2

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result = &ListNode{}
	carry := 0
	tail := result;
	for {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		tail.Val = sum % 10
		carry = sum / 10
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		tail.Next = &ListNode{}
		tail = tail.Next
	}
	return result
}

package q19

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var pos = make(map[int]*ListNode, n+1)
	var length = 0
	for p := head; p != nil; p = p.Next {
		pos[length] = p
		length++
		if length-n-2 >= 0 {
			delete(pos, length-n-2)
		}
	}
	toDelIdx := length - n

	if toDelIdx == 0 {
		head = head.Next
	} else {
		pos[toDelIdx-1].Next = pos[toDelIdx+1]
	}
	return head
}

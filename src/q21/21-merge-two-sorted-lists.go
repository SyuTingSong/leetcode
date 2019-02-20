package q21

import "fmt"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var result = &ListNode{}
	var tail = result
	for l1 != nil && l2 != nil {
		node := &ListNode{}
		if l1.Val < l2.Val {
			node.Val = l1.Val
			l1 = l1.Next
		} else {
			node.Val = l2.Val
			l2 = l2.Next
		}
		tail.Next = node
		tail = tail.Next
	}
	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}
	return result.Next
}

func main() {
	l1 := FromIntegers([]int{1, 2, 4})
	l2 := FromIntegers([]int{1, 3, 4})
	r := mergeTwoLists(l1, l2)
	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println(r)
}

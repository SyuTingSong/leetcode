package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func main() {
	r := addTwoNumbers(
		&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		&ListNode{4, &ListNode{5, &ListNode{6, nil}}},
	)

	for n := r; n != nil; n = n.Next {
		fmt.Println(n.Val)
	}
}

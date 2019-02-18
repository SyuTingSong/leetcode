package main

import "fmt"

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
	n1 := FromIntegers([]int{1, 2, 3})
	n2 := FromIntegers([]int{4, 5, 6})
	r := addTwoNumbers(n1, n2)

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(r)
}

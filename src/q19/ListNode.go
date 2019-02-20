package q19

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func FromIntegers(integers []int) *ListNode {
	l := len(integers)
	root := &ListNode{}
	p := root
	for i := 0; i < l; i++ {
		p.Next = &ListNode{
			integers[i],
			nil,
		}
		p = p.Next
	}
	return root.Next
}

func (l *ListNode) Format(s fmt.State, c rune) {
	b := &strings.Builder{}
	for p := l; p != nil; p = p.Next {
		b.WriteString(fmt.Sprintf("%d -> ", p.Val))
	}
	b.WriteString("nil")
	_, _ = s.Write([]byte(b.String()))
}

func (l *ListNode) Same(t *ListNode) bool {
	var p, q *ListNode
	for p, q = l, t; p != nil && q != nil; p, q = p.Next, q.Next {
		if p.Val != q.Val {
			return false
		}
	}
	return q == nil && p == nil
}

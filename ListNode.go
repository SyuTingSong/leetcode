package main

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

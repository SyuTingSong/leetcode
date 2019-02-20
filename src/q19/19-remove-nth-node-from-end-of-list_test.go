package q19

import (
	"testing"
)

func listNodePoints(node *ListNode) []*ListNode {
	var p []*ListNode
	for n := node; n != nil; n = n.Next {
		p = append(p, node)
	}
	return p
}

func match(node *ListNode, nodes []*ListNode) bool {
	for n := 0; n < len(nodes); n++ {
		if nodes[n] != node {
			return false
		}
		node = node.Next
	}
	return node.Next == nil
}

func Test_removeNthFromEnd(t *testing.T) {
	list := FromIntegers([]int{1, 2, 3, 4, 5})
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"[1, 2, 3, 4, 5], 2",
			args{
				list,
				2,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pointers := listNodePoints(tt.args.head)
			n := len(pointers)
			x := append(pointers[:n-tt.args.n], pointers[n-tt.args.n+1:]...)
			got := removeNthFromEnd(tt.args.head, tt.args.n)
			if !match(got, x) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, x[0])
			}
		})
	}
}

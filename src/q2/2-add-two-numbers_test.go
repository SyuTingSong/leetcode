package q2

import "testing"

type testCase struct {
	l1       *ListNode
	l2       *ListNode
	expected *ListNode
}

var cases = []testCase{
	{
		FromIntegers([]int{2, 4, 3}),
		FromIntegers([]int{5, 6, 4}),
		FromIntegers([]int{7, 0, 8}),
	},
	{
		FromIntegers([]int{2, 4, 3}),
		FromIntegers([]int{5, 6, 5}),
		FromIntegers([]int{7, 0, 9}),
	},
}

func TestAddTwoNumbers(t *testing.T) {
	for _, c := range cases {
		r := addTwoNumbers(c.l1, c.l2)
		if !r.Same(c.expected) {
			t.Errorf(
				"Expected %v\ngot %v\nl1: %v\nl2: %v",
				c.expected, r, c.l1, c.l2,
			)
		}
	}
}

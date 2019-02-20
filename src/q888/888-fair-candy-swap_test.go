package q888

import (
	"testUtil"
	"testing"
)

type testCase struct {
	A        []int
	B        []int
	expected []int
}

var cases = []testCase{
	{
		[]int{1, 3},
		[]int{2},
		[]int{3, 2},
	},
}

func TestFairCandySwap(t *testing.T) {
	for _, c := range cases {
		r := fairCandySwap(c.A, c.B)
		if !testUtil.IntSlice(r).Equal(c.expected) {
			t.Errorf(
				"Expected %v got %v when A: %v B: %v",
				c.expected, r, c.A, c.B,
			)
		}
	}
}

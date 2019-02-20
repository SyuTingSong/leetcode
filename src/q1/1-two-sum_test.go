package q1

import (
	"testUtil"
	"testing"
)

type testCase struct {
	nums     []int
	target   int
	expected []int
}

var cases = []testCase{
	{
		[]int{2, 7, 11, 15},
		9,
		[]int{0, 1},
	},
}

func TestTwoSum(t *testing.T) {
	for _, c := range cases {
		r := twoSum(c.nums, c.target)
		if !testUtil.IntSlice(r).Same(c.expected) {
			t.Errorf("Expected %v got %v nums: %v target %v", c.expected, r, c.nums, c.target)
		}
	}
}

package q4

import "testing"

type testCase struct {
	nums1    []int
	nums2    []int
	expected float64
}

var cases = []testCase{
	{[]int{1, 3}, []int{2}, 2},
	{[]int{1, 2}, []int{3, 4}, 2.5},
}

func TestFindMedianSortedArrays(t *testing.T) {
	for _, c := range cases {
		r := findMedianSortedArrays(c.nums1, c.nums2)
		if r != c.expected {
			t.Errorf("Expected %v got %v when %v %v", c.expected, r, c.nums1, c.nums2)
		}
	}
}

package q3

import "testing"

type testCase struct {
	s        string
	expected int
}

var cases = []testCase{
	{"abcabcbb", 3},
	{"bbbbb", 1},
	{"pwwkew", 3},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, c := range cases {
		r := lengthOfLongestSubstring(c.s)
		if r != c.expected {
			t.Errorf("Expected %v got %v when %s", c.expected, r, c.s)
		}
	}
}

package q950

import (
	"testUtil"
	"testing"
)

type testCase struct {
	deck     []int
	expected []int
}

var cases = []testCase{
	{
		[]int{17, 13, 11, 2, 3, 5, 7},
		[]int{2, 13, 3, 11, 5, 17, 7},
	},
}

func TestDeckRevealedIncreasing(t *testing.T) {
	for _, c := range cases {
		r := deckRevealedIncreasing(c.deck)
		if !testUtil.IntSlice(r).Equal(c.expected) {
			t.Errorf("Expect %v got %v, deck %v", c.expected, r, c.deck)
		}
	}
}

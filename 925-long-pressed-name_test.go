package main

import "testing"

type testSuite struct {
	name     string
	typed    string
	expected bool
}

var tests = []testSuite{
	{
		"alex",
		"aaleex",
		true,
	},
	{
		"saeed",
		"ssaaedd",
		false,
	},
	{
		"leelee",
		"lleeelee",
		true,
	},
}

func Test925(t *testing.T) {
	for _, c := range tests {
		r := isLongPressedName(c.name, c.typed)
		if r == c.expected {
			t.Log("OK")
		} else {
			t.Errorf("Expected %v got %v, name: %s, typed: %s\n", c.expected, r, c.name, c.typed)
		}
	}
}

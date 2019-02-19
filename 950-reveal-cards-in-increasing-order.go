package main

import (
	"container/list"
	"sort"
)

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)
	l := list.New()
	for n := len(deck) - 1; n >= 0; n-- {
		if back := l.Back(); back != nil {
			l.MoveToFront(back)
		}
		card := deck[n]
		l.PushFront(card)
	}
	var result []int
	for e := l.Front(); e != nil; e = e.Next() {
		if i, ok := e.Value.(int); ok {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	deckRevealedIncreasing([]int{17, 13, 11, 2, 3, 5, 7})
}

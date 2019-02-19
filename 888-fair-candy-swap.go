package main

import "fmt"

func sum(c []int) (s int) {
	l := len(c)
	for i := 0; i < l; i++ {
		s += c[i]
	}
	return
}

func fairCandySwap(A []int, B []int) []int {
	var result []int
	b2a := make(map[int]int)
	diff := (sum(A) - sum(B)) / 2
	for _, a := range A {
		b2a[a-diff] = a
	}
	for _, b := range B {
		if a, ok := b2a[b]; ok {
			result = []int{a, b}
			break
		}
	}
	return result
}

func main() {
	fmt.Println(fairCandySwap([]int{1, 3}, []int{2}))
}

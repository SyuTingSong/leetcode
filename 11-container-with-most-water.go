package main

import "fmt"

func maxArea(height []int) int {
	s := 0
	e := len(height) - 1
	max := 0
	for s < e {
		area := 0
		if height[s] < height[e] {
			area = height[s] * (e - s)
			s++
		} else {
			area = height[e] * (e - s)
			e--
		}
		if area > max {
			max = area
		}
	}
	return max
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

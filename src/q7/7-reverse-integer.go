package q7

import (
	"math"
)

func safeReverse(x int) (r int) {
	for x != 0 {
		r = r*10 + x%10
		x /= 10
	}
	return
}
func reverse(x int) int {
	var negative = x < 0
	var r int
	if negative {
		x = -x
	}
	var size = int(math.Floor(math.Log10(float64(x))))
	if size < 10 {
		r = safeReverse(x)
	} else {
		return 0
	}
	if negative {
		return -r
	}
	return r
}

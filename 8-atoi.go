package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(str string) int {
	var num = 0
	s := strings.TrimLeft(str, "\r\n\t ")[:]
	negative := false
	for i, c := range s {
		if i == 0 {
			if c == '+' || c == '-' {
				negative = c == '-'
			} else if c >= '0' && c <= '9' {
				num = num*10 + int(c&0x0f)
			} else {
				return 0
			}
		} else {
			if c >= '0' && c <= '9' {
				if num > math.MaxInt32/10 {
					if negative {
						return math.MinInt32
					} else {
						return math.MaxInt32
					}
				} else if num == math.MaxInt32/10 {
					if negative {
						if c > '8' {
							return math.MinInt32
						}
					} else {
						if c > '7' {
							return math.MaxInt32
						}
					}
				}
				num = num*10 + int(c&0x0f)
			} else {
				break
			}
		}
	}
	if negative {
		return -num;
	}
	return num
}

func main() {
	fmt.Println(myAtoi("42"))
}

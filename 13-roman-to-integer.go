package main

import (
	"fmt"
)

func romanToInt(s string) int {
	var result int
	chars := s[:]
	length := len(chars)
	var next uint8
	for i := 0; i < length; i++ {
		c := chars[i]
		if i+1 < length {
			next = chars[i+1]
		} else {
			next = 0
		}
		switch c {
		case 'M':
			result += 1000
		case 'D':
			result += 500
		case 'C':
			if next == 'M' || next == 'D' {
				result -= 100
			} else {
				result += 100
			}
		case 'L':
			result += 50
		case 'X':
			if next == 'C' || next == 'L' {
				result -= 10
			} else {
				result += 10
			}
		case 'V':
			result += 5
		case 'I':
			if next == 'X' || next == 'V' {
				result -= 1
			} else {
				result += 1
			}
		}
	}
	return result
}

func main() {
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

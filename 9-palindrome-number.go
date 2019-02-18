package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x > 0 && x%10 == 0 {
		return false
	}

	var upper = x
	var lower = 0
	for upper > lower {
		lower = lower*10 + upper%10
		upper /= 10
	}
	return upper == lower || upper == lower/10
}

func main() {
	fmt.Println(isPalindrome(121))
}

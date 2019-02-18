package main

import "fmt"

func isValid(s string) bool {
	if s == "" {
		return true
	}
	var parentheses []uint8
	l := len(s)
	n := 0
	for i := 0; i < l; i++ {
		//fmt.Println(n, parentheses)
		c := s[i]
		if c == '(' || c == '[' || c == '{' {
			parentheses = append(parentheses, c)
			n++
		} else {
			if n == 0 {
				return false
			}
			if c == ')' && parentheses[n-1] != '(' {
				return false
			}
			if c == ']' && parentheses[n-1] != '[' {
				return false
			}
			if c == '}' && parentheses[n-1] != '{' {
				return false
			}

			n--
			if n > 0 {
				parentheses = parentheses[:n]
			} else {
				parentheses = nil
			}
		}
		//fmt.Println(n, parentheses)
		//fmt.Println("----")
	}
	return n == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("([])"))
	fmt.Println(isValid("([()(){}])"))
	fmt.Println(isValid("([()(){}]"))
	fmt.Println(isValid(")"))
}

package src

import (
	"container/list"
	"fmt"
)

func generateParenthesis(n int) []string {
	r := list.New()
	parenthesisTry(r, "", 0, 0, n)
	var result []string
	for s := r.Front(); s != nil; s = s.Next() {
		if str, ok := s.Value.(string); ok {
			result = append(result, str)
		}
	}
	return result
}

func parenthesisTry(r *list.List, p string, open int, close int, max int) {
	if len(p) == max*2 {
		r.PushBack(p)
		return
	}

	if open < max {
		parenthesisTry(r, p+"(", open+1, close, max)
	}
	if close < open {
		parenthesisTry(r, p+")", open, close+1, max)
	}
}

func main() {
	fmt.Println(generateParenthesis(5))
}

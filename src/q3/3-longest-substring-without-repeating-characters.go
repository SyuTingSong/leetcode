package q3

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var chars = s[:]
	var length = len(chars)
	var charSet = make(map[uint8]int)
	if length <= 1 {
		return length
	} else if length == 2 {
		if chars[0] == chars[1] {
			return 1
		} else {
			return 2
		}
	}

	head, tail, max, cur := 0, 0, 0, 0
	for tail < length {
		char := chars[tail]
		//fmt.Println(head, tail, charSet, char)
		if p, ok := charSet[char]; ok {
			if cur > max {
				max = cur
			}
			for head <= p {
				delete(charSet, chars[head])
				head++
			}
			charSet[char] = tail
			if tail < head {
				tail = head
			}
			tail++
			cur = tail - head
		} else {
			charSet[char] = tail
			cur++
			tail++
			if cur > max {
				max = cur
			}
		}
		//fmt.Println(head, tail, charSet, cur, max)
		//fmt.Println("------")
	}
	if cur > max {
		return cur
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcdefbxyzqts"))
}

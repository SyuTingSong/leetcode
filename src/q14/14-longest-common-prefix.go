package q14

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var numOfStr = len(strs)
	if numOfStr == 0 {
		return ""
	}
	if numOfStr == 1 {
		return strs[0]
	}

	var minLength = len(strs[0])
	for i := 1; i < numOfStr; i++ {
		if len(strs[i]) < minLength {
			minLength = len(strs[i])
		}
	}

	var common = &strings.Builder{}
	for p := 0; p < minLength; p++ {
		var char = strs[0][p]
		for i := 1; i < numOfStr; i++ {
			if strs[i][p] != char {
				return common.String()
			}
		}
		common.WriteByte(char)
	}
	return common.String()
}

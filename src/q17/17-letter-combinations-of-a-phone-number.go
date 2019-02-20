package q17

import (
	"strings"
)

func letterCombinations(digits string) []string {
	dict := map[uint8]string{
		'2': "abc"[:],
		'3': "def"[:],
		'4': "ghi"[:],
		'5': "jkl"[:],
		'6': "mno"[:],
		'7': "pqrs"[:],
		'8': "tuv"[:],
		'9': "wxyz"[:],
	}
	digitArr := digits[:]
	length := len(digitArr)
	if length == 0 {
		return nil
	}
	num := 1
	for n := length; n > 0; n-- {
		digit := digitArr[n-1]
		num *= len(dict[digit])
	}

	var result []string
	for i := 0; i < num; i++ {
		t := i
		b := &strings.Builder{}
		for n := length; n > 0; n-- {
			digit := digitArr[length-n]
			base := len(dict[digit])
			b.WriteByte(dict[digit][t%base])
			t /= base
		}
		result = append(result, b.String())
	}

	return result
}

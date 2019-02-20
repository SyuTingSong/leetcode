package q6

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var rows = make([]*strings.Builder, numRows)
	for i := 0; i < numRows; i++ {
		rows[i] = &strings.Builder{}
	}

	var idx = 0
	var length = len(s)
	var step = 1

	for i := 0; i < length; i++ {
		//fmt.Println(i, idx)
		rows[idx].WriteByte(s[i])
		if idx == 0 {
			step = 1
		} else if idx == numRows-1 {
			step = -1
		}
		idx += step
	}
	var b = &strings.Builder{}
	for i := 0; i < numRows; i++ {
		b.WriteString(rows[i].String())
	}
	return b.String()
}

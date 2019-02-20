package q12

type M struct {
	ten   int
	roman string
}

func intToRoman(num int) string {
	var units = []M{
		M{1000, "M"},
		M{900, "CM"},
		M{500, "D"},
		M{400, "CD"},
		M{100, "C"},
		M{90, "XC"},
		M{50, "L"},
		M{40, "XL"},
		M{10, "X"},
		M{9, "IX"},
		M{5, "V"},
		M{4, "IV"},
		M{1, "I"},
	}

	var result = ""
	for _, u := range units {
		for n := num / u.ten; n > 0; n-- {
			result += u.roman
		}
		num = num % u.ten
	}
	return result
}

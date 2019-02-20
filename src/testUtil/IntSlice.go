package testUtil

type IntSlice []int

func (s IntSlice) Equal(b []int) bool {
	for i, _ := range s {
		if s[i] != b[i] {
			return false
		}
	}
	return true
}

func (s IntSlice) Same(b []int) bool {
	set := make(map[int]bool)
	for _, i := range s {
		set[i] = false
	}
	for _, i := range b {
		if _, ok := set[i]; ok {
			set[i] = true
		} else {
			return false
		}
	}
	for _, hit := range set {
		if !hit {
			return false
		}
	}
	return true
}

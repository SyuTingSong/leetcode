package testUtil

type StringSlice []string

func (s StringSlice) Same(b []string) bool {
	set := make(map[string]bool)
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

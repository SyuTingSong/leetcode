package q925

func isLongPressedName(name string, typed string) bool {
	ln, lt := len(name), len(typed)
	i, j, hit := 0, 0, false
	for j < lt {
		if name[i] == typed[j] {
			if i+1 < ln && name[i+1] == name[i] {
				i++
				j++
			} else {
				hit = true
				j++
			}
		} else if i < ln && hit {
			hit = false
			i++
		} else {
			return false
		}
	}
	return j == lt && i == ln-1 && hit
}

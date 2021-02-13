package feb

func isAnagram(s string, t string) bool {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}

	for _, r := range t {
		if m[r] == 0 {
			return false
		}

		m[r] -= 1
		if m[r] == 0 {
			delete(m, r)
		}
	}

	return len(m) == 0
}

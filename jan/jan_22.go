package leetcode

func mapper(word string) map[rune]int {
	result := make(map[rune]int)

	for _, r := range word {
		result[r]++
	}

	return result
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	w1 := mapper(word1)
	w2 := mapper(word2)

	for r, v := range w2 {
		if w1[r] == 0 {
			return false
		}

		if w1[r] == v {
			continue
		}

		for k, c := range w1 {
			if c == v {
				w1[k] = w1[r]
				w1[r] = c
				break
			}
		}

		return false
	}

	return true
}

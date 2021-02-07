package feb

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func shortestToChar(s string, c byte) []int {
	occures := make([]int, 0)
	for i, e := range s {
		if rune(c) == e {
			occures = append(occures, i)
		}
	}

	result := make([]int, len(s))
	minIdx := 0
	for i := range s {
		if occures[minIdx] == i {
			result[i] = 0
			minIdx++
		} else {
			before := int(1e5)
			after := int(1e5)

			if minIdx < len(occures) {
				before = occures[minIdx] - i
			}

			if minIdx-1 > 0 {
				after = i - occures[minIdx-1]
			}

			result[i] = min(before, after)
		}
	}

	return result
}

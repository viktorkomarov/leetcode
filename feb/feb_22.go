package feb

func buildIndexMap(s string) map[rune][]int {
	result := make(map[rune][]int, 0)
	for i, r := range s {
		result[r] = append(result[r], i)
	}

	return result
}

func canBuild(dic map[rune][]int, str string) bool {
	after := -1
	for _, r := range str {
		indx, ok := dic[r]
		if !ok {
			return false
		}

		keep := after
		for _, j := range indx {
			if j > after {
				after = j
				break
			}
		}

		if keep == after {
			return false
		}
	}

	return true
}

func findLongestWord(s string, d []string) string {
	dic := buildIndexMap(s)

	max := ""
	for _, str := range d {
		if canBuild(dic, str) {
			if len(str) > len(max) || (len(str) == len(max) && str < max) {
				max = str
			}
		}
	}

	return max
}

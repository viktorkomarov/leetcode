package feb

func removeInvalidBack(runes []rune) []rune {
	clean := make([]rune, 0, len(runes))

	countOpen := 0
	for _, r := range runes {
		add := true

		if r == '(' {
			countOpen++
		}

		if r == ')' {
			if countOpen > 0 {
				countOpen--
			} else {
				add = false
			}
		}

		if add {
			clean = append(clean, r)
		}
	}

	return clean
}

func removeInvalidOpen(runes []rune) []rune {
	countClose := 0
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == ')' {
			countClose++
		}

		if runes[i] == '(' {
			if countClose > 0 {
				countClose--
			} else {
				runes[i] = '\000'
			}
		}
	}

	clean := make([]rune, 0, len(runes))
	for i := range runes {
		if runes[i] != '\000' {
			clean = append(clean, runes[i])
		}
	}
	return clean
}

func minRemoveToMakeValid(s string) string {
	runes := []rune(s)
	return string(removeInvalidOpen(removeInvalidBack(runes)))
}

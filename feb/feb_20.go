package feb

func numByRune(r rune) int {
	switch r {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}

	return 0
}

func romanToInt(s string) int {
	result := 0

	prev := -1
	for _, r := range s {
		num := numByRune(r)
		if num > prev && prev != -1 {
			result += num - 2*prev
		} else {
			result += num
		}

		prev = num
	}

	return result
}

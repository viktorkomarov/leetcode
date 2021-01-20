package leetcode

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))

	not_reverse := func(c, t rune) bool {
		if c == ')' {
			return t != '('
		}
		if c == '}' {
			return t != '{'
		}

		return t != '['
	}

	for _, r := range s {
		switch r {
		case '(', '[', '{':
			stack = append(stack, r)
		case ']', '}', ')':
			if len(stack) == 0 || not_reverse(r, stack[len(stack)-1]) {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

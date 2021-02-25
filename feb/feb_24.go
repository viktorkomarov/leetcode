package feb

func scoreOfParentheses(S string) int {
	ans := 0
	bal := 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			bal++
		} else {
			bal--
			if S[i-1] == '(' {
				ans += 1 << bal
			}
		}
	}

	return ans
}

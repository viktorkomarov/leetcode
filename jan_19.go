package leetcode

func longestPalindrome(s string) string {
	table := make([][]bool, len(s))
	for i := range table {
		table[i] = make([]bool, len(s))
	}

	max := 1
	start := 0
	for i := 0; i < len(s); i++ {
		table[i][i] = true
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			max = 2
			start = i
			table[i][i+1] = true
		}
	}

	for k := 3; k <= len(s); k++ {
		for i := 0; i < len(s)-k+1; i++ {
			j := i + k - 1

			if table[i+1][j-1] && s[i] == s[j] {
				table[i][j] = true

				if k > max {
					max = k
					start = i
				}

			}
		}
	}

	return s[start : start+max]
}

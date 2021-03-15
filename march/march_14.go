package march

func isSubsequence(s string, t string) bool {
	id := 0

	for i := range t {
		if id >= len(s) {
			break
		}

		if t[i] == s[id] {
			id++
		}
	}

	return id == len(s)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func countSquares(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i := range matrix {
		dp[i] = make([]int, len(matrix[i]))
	}
	
	res := 0

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 1 {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
                    dp[i][j] = min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
				}

				res += dp[i][j]
			}
		}
	}

	return res
}
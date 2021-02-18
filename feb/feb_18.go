package feb

func numberOfArithmeticSlices(A []int) int {
	dp := make([]int, len(A))
	sum := 0

	for i := 2; i < len(A); i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			dp[i] += 1 + dp[i-1]
		}
		sum += dp[i]
	}

	return sum
}

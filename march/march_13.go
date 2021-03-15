package march

func divisorGame(N int) bool {
	nums := make([]int, N+1)
	
	for i := 2; i <= N; i++ {
		nums[i] = nums[i-1]+1
	}

	return nums[N] % 2 != 0
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minCostClimbingStairs(cost []int) int {
    dp := make([]int, len(cost)+1)

	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}

    return min(dp[len(cost)-1],dp[len(cost)-2])
}

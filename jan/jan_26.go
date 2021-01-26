package jan

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(val int) int {
	if val < 0 {
		return -val
	}

	return val
}

func minimumEffortPath(heights [][]int) int {
	result := make([][]int, len(heights))
	for i := range result {
		result[i] = make([]int, len(heights[i]))
		for j := range result[i] {
			result[i][j] = 1_000_001
		}
	}
	result[0][0] = 0

	toLeft := func(i, j int) {
		if j-1 >= 0 {
			result[i][j-1] = min(result[i][j], min(result[i][j-1], abs(heights[i][j]-heights[i][j-1])))
		}
	}

	toRight := func(i, j int) {
		if j+1 < len(heights[i]) {
			result[i][j+1] = min(result[i][j], min(result[i][j+1], abs(heights[i][j]-heights[i][j+1])))
		}
	}

	toUp := func(i, j int) {
		if i+1 < len(heights) {
			result[i+1][j] = min(result[i][j], min(result[i+1][j], abs(heights[i][j]-heights[i+1][j])))
		}
	}

	toBottom := func(i, j int) {
		if i-1 >= 0 {
			result[i-1][j] = min(result[i][j], min(result[i-1][j], abs(heights[i][j]-heights[i-1][j])))
		}
	}

	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[i]); j++ {
			toUp(i, j)
			toBottom(i, j)
			toLeft(i, j)
			toRight(i, j)
		}
	}

	l := result[len(result)-1]
	return l[len(l)-1]
}

package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func lengthOfLongestSubstring(s string) int {
	result := 0

	symbols := make([]int, 256)
	for i := range symbols {
		symbols[i] = -1
	}

	startSubstring := 0
	for i, r := range s {
		startSubstring = max(startSubstring, symbols[int(r)]+1)
		result = max(result, i-startSubstring+1)
		symbols[int(r)] = i
	}

	return result
}

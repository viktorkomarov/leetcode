package leetcode

func findKthPositive(arr []int, k int) int {
	currentPositive := 1
	for _, in := range arr {
		gap := in - currentPositive
		if gap > 0 {
			if k <= gap {
				return currentPositive + k - 1
			} else {
				k -= gap
			}
		}

		currentPositive = in + 1
	}

	return currentPositive + k - 1
}

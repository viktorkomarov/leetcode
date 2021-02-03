package arr

func findMaxConsecutiveOnes(nums []int) int {
	result := 0

	current := 0
	for _, n := range nums {
		if n == 1 {
			current++
		} else {
			if current > result {
				result = current
			}
			current = 0
		}
	}

	if current > result {
		result = current
	}

	return result
}

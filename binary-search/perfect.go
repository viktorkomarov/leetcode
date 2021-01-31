package binary_search

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	left := 0
	right := num

	for left < right-1 {
		mid := left + (right-left)/2
		if mid*mid == num {
			return true
		}

		if mid*mid > num {
			right = mid
		} else {
			left = mid
		}
	}

	return false
}

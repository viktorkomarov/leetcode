package binary_search

func mySqrt(x int) int {
	left := 0
	right := x / 2

	for left < right {
		mid := left + ((right - left) / 2)

		if mid*mid > x {
			right = mid
		} else {
			left = mid
		}
	}

	return left
}

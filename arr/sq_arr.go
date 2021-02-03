package arr

func sortedSquares(nums []int) []int {
	var (
		min = int(1e5)
		max = int(-1e5)
	)
	counts := make(map[int]int)
	for _, n := range nums {
		if n < 0 {
			n = -n
		}

		counts[n]++
		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	i := 0
	for min <= max {
		for ; counts[min] > 0; counts[min]-- {
			nums[i] = min * min
			i++
		}
		min++
	}

	return nums
}

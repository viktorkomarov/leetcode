package march

func missingNumber(nums []int) int {
	sum := (len(nums) + 1) * len(nums) / 2

	current := 0
	for _, n := range nums {
		current += n
	}

	return sum - current
}

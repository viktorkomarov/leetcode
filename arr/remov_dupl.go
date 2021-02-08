package arr

func removeDuplicates(nums []int) int {
	available := 0
	prev := int(1e5)

	for _, v := range nums {
		if v != prev {
			nums[available] = v
			prev = v
			available++
		}
	}

	return available
}

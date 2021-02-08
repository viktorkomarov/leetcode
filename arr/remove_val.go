package arr

func removeElement(nums []int, val int) int {
	count := 0

	for j, num := range nums {
		if num == val {
			count++
			continue
		}

		idx := j
		for i := 0; i < count && idx-1 >= 0; i++ {
			nums[idx-1] = nums[idx]
			idx--
		}

	}

	return count
}

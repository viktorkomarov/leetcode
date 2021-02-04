package feb

func findLHS(nums []int) int {
	levels := make(map[int]int)

	for _, num := range nums {
		levels[num]++
	}

	max := 0
	for key, val := range levels {
		before := levels[key-1]
		after := levels[key+1]

		if before != 0 && val+before > max {
			max = val + before
		}

		if after != 0 && val+after > max {
			max = val + before
		}
	}

	return max
}

package march

func findErrorNums(nums []int) []int {
	num := make(map[int]int)

	for i := range nums {
		num[nums[i]]++
	}

	o, d := 0, 1
	for i := 1; i < int(1e4); i++ {
		c, ok := num[i]
		if !ok && o == 0 {
			o = i
		}

		if c == 2 {
			d = i
		}

	}

	return []int{d, o}
}

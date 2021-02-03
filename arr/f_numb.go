package arr

func findNumbers(nums []int) int {
	digitNum := func(n int) int {
		count := 0
		for n > 0 {
			count++
			n /= 10
		}

		return count
	}
	even := 0
	for _, n := range nums {
		if digitNum(n)%2 == 0 {
			even++
		}
	}
	return even
}

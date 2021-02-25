package feb

import "sort"

func findUnsortedSubarray(nums []int) int {
	c := make([]int, len(nums))
	copy(c, nums)

	sort.Ints(c)

	l := 0
	for l < len(nums) && nums[l] == c[l] {
		l++
	}
	if l == len(nums) {
		return 0
	}

	r := len(nums) - 1
	for r > l && nums[r] == c[r] {
		r--
	}

	return r - l + 1
}

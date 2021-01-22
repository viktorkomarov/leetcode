package leetcode

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)

	left := 0
	right := len(nums) - 1
	result := 0

	for left < right {
		sum := nums[left] + nums[right]
		if sum == k {
			result++
			left++
			right--
		} else if sum < k {
			left++
		} else {
			right--
		}
	}

	return result
}

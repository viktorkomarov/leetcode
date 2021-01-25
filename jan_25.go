package leetcode

func kLengthApart(nums []int, k int) bool {
	prev := -1

	for i := range nums {
		if nums[i] == 1 {
			if prev != -1 {
				if i-prev < k {
					return false
				}
			}
			prev = i
		}
	}

	return true
}

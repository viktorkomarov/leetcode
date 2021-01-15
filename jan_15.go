package leetcode

func getMaximumGenerated(n int) int {
	if == 0 {
		return 0
	}
	if n == 1{
		return 1
	}

	nums := make([]int, n)
	nums[1] = 1

	for i := 2; i < len(nums); i++ {
		if i%2 == 0 {
			nums[i] = nums[i-1]
		} else {
			nums[i] = nums[i-1] + nums[i-2]
		}
	}

	if num[len(nums)-1] > num[len(nums)-2] {
		return num[len(nums)-1]
	}

	return num[len(nums)-2]
}

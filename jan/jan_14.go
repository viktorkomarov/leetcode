package leetcode

func minOperations(nums []int, x int) int {
	leftAcc := 0
	rightAcc := 0

	for i := 0; i < len(nums); i++ {
		left := nums[i]
		right := nums[len(nums)-1-i]

		if leftAcc+left == x || rightAcc+right == x {
			return i + 1
		}

		rights := rightAcc
		lefts := leftAcc
		for j := 0; j < i; j++ {
			lefts += nums[len(nums)-j-1]
			if lefts == x {
				return i + j + 1
			}

			rights += nums[j]
			if rights == x {
				return i + j + 1
			}
		}

		leftAcc += left
		rightAcc += right
	}

	return -1
}

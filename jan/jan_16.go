package leetcode

func partioner(nums []int, l, r int) int {
	mid := nums[(l+r)/2]
	for l <= r {
		for nums[l] < mid {
			l++
		}

		for nums[r] > mid {
			r--
		}

		if l >= r {
			break
		}

		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}

	return r
}

func findKthLargest(nums []int, k int) int {
	l := 0
	r := len(nums) - 1
	targer := len(nums) - k
	for {
		n := partioner(nums, l, r)
		if n == targer {
			return nums[n]
		}

		if n > targer {
			r = n - 1
		} else {
			l = n
		}

	}
}

/*

func findKthLargest(nums []int, k int) int {
    sort.Ints(nums)

    return nums[len(nums)-k]
}

*/

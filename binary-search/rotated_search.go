package binary_search

func min(nums []int) int {
	left := 0
	right := len(nums) - 1

	for {
		if left == right {
			return left
		}

		mid := left + ((right - left) / 2)

		if nums[mid] > nums[mid+1] {
			return mid + 1
		}

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
}

func search(nums []int, target int) int {
	n := len(nums)
	pivot := min(nums)
	left, right := pivot, pivot-1+n
	for left <= right {
		mid := left + (right-left)/2
		midVal := nums[mid%n]

		if midVal > target {
			right = mid - 1
		} else if midVal < target {
			left = mid + 1
		} else {
			return mid % n
		}
	}

	return -1
}

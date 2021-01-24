package binary_search

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for {
		if left >= right {
			return nums[left]
		}

		mid := left + ((right - left) / 2)

		if mid > 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		}

		if mid < len(nums) && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}

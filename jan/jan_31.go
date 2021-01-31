package jan

func nextPermutation(nums []int) {
	broken := len(nums) - 2
	for broken >= 0 && nums[broken+1] <= nums[broken] {
		broken--
	}

	if broken >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[broken] {
			j--
		}

		nums[broken], nums[j] = nums[j], nums[broken]
	}

	reverse(nums, broken+1)
}

func reverse(nums []int, from int) {
	strart := from
	end := len(nums) - 1

	for strart < end {
		nums[strart], nums[end] = nums[end], nums[strart]
		strart++
		end--
	}
}

package jan

import (
	"fmt"
	"sort"
)

// wrong :с
func getMin(i int) int {
	for i%2 == 0 {
		i /= 2
	}

	return i
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func minimumDeviation(nums []int) int {
	containTwo := false
	for i, num := range nums {
		if num == 2 {
			containTwo = true
			continue
		}

		nums[i] = getMin(num)
	}

	max := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == 2 || nums[i] == max {
			continue
		}

		if nums[i] == 1 {
			nums[i] = 2
			containTwo = true
			continue
		}

		if containTwo {
			if (nums[i]*2 - 2 + abs(max-nums[i]*2)) < (nums[i] - 2 + max - nums[i]*2) {
				nums[i] *= 2
			}
		} else {
			if abs(max-nums[i]*2) < abs(max-nums[i]) {
				nums[i] *= 2
			}
		}

	}

	fmt.Printf("%+v\n", nums)
	sort.Ints(nums)
	return nums[len(nums)-1] - nums[0]
}

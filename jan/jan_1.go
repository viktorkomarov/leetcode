package leetcode

func canFormArray(arr []int, pieces [][]int) bool {
	result := make([]int, 0, len(arr))
	for _, val := range arr {
		for _, part := range pieces {
			if val == part[0] {
				result = append(result, part...)
				break
			}
		}
	}

	if len(result) == len(arr) {
		for i := range arr {
			if arr[i] != result[i] {
				return false
			}
		}

		return true
	}

	return false
}

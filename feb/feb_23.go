package feb

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	for i := range matrix {
		if idx := sort.Search(len(matrix[i]), func(n int) bool {
			return matrix[i][n] >= target
		}); idx < len(matrix[i]) && matrix[i][idx] == target {
			return true
		}
	}

	return false
}

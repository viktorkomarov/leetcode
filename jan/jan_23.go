package leetcode

func diagonalSort(mat [][]int) [][]int {
	for i := 1; i < len(mat); i++ {
		for j := 1; j < len(mat[i]); j++ {
			line := i - 1
			column := j - 1
			for line >= 0 && column >= 0 && mat[line][column] > mat[line+1][column+1] {
				mat[line][column], mat[line+1][column+1] = mat[line+1][column+1], mat[line][column]
				line--
				column--
			}
		}
	}
	return mat
}

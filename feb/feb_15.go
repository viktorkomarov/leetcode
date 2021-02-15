package feb

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	type stat struct {
		idx   int
		count int
	}
	stats := make([]stat, 0)

	for i := range mat {
		zeroNum := 0
		for j := range mat[i] {
			if mat[i][j] == 0 {
				zeroNum++
			}
		}
		stats = append(stats, stat{
			idx:   i,
			count: zeroNum,
		})
	}

	sort.Slice(stats, func(i, j int) bool {
		if stats[i].count == stats[j].count {
			return stats[i].idx < stats[j].idx
		}

		return stats[i].count > stats[j].count
	})

	result := make([]int, k)
	for i := range result {
		result[i] = stats[i].idx
	}

	return result
}

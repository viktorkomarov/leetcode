package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func createSortedArray(instructions []int) int {
	cost := 0
	for i := 1; i < len(instructions); i++ {
		num := instructions[i]
		j := i - 1

		counter := 0
		for ; j >= 0 && num <= instructions[j]; j-- {
			instructions[j], instructions[j+1] = num, instructions[j]
			if num != instructions[j+1] {
				counter++
			}
		}
		cost += min(j+1, counter)
	}
	return cost
}

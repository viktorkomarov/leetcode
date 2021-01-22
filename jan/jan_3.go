package leetcode

func generate(use [][]int, index int, stopList map[int]bool) int {
	if index == len(use) {
		return 1
	}

	end := 0
	for _, in := range use[index] {
		if !stopList[in] {
			stopList[in] = true
			end += generate(use, index+1, stopList)
			stopList[in] = false
		}
	}

	return end
}

func countArrangement(n int) int {
	useByIndex := make([][]int, n)
	for i := range useByIndex {
		for j := 1; j <= n; j++ {
			if (i+1)%j == 0 || j%(i+1) == 0 {
				useByIndex[i] = append(useByIndex[i], j)
			}
		}
	}

	return generate(useByIndex, 0, map[int]bool{})
}

package march

func distributeCandies(candyType []int) int {
	counts := make(map[int]int)

	for _, n := range candyType {
		counts[n]++
	}

	if len(counts) < len(candyType)/2 {
		return len(counts)
	}

	return len(candyType) / 2
}

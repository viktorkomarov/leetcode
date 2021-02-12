package other

func batch(count int, people []int) [][]int {
	batched := make([][]int, 0)

	for i := 0; i < len(people); i += count {
		batched = append(batched, people[i:i+count])
	}

	return batched
}

func groupThePeople(groupSizes []int) [][]int {
	groups := make(map[int][]int)
	for i, key := range groupSizes {
		groups[key] = append(groups[key], i)
	}

	result := make([][]int, 0)
	for count, people := range groups {
		result = append(result, batch(count, people)...)
	}

	return result
}

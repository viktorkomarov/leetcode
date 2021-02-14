package feb

type color int

const (
	Unknown color = iota
	Red
	White
)

func change(c color) color {
	if c == Red {
		return White
	}

	return Red
}

func isBipartite(graph [][]int) bool {
	parts := make(map[int]color)

	queue := make([]int, 0)
	current := Red
	for i := range graph {
		if _, ok := parts[i]; ok {
			continue
		}

		queue = append(queue, i)
		for len(queue) != 0 {
			for _, v := range queue {
				if _, ok := parts[v]; !ok {
					parts[v] = current
				}

				current = parts[v]
				for _, ni := range graph[v] {
					part, ok := parts[ni]
					if !ok {
						parts[ni] = change(current)
						queue = append(queue, ni)
						continue
					}
					if part == current {
						return false
					}
				}

				current = change(current)
				queue = queue[1:]
			}
		}
	}

	return true
}

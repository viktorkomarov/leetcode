package feb

type Node struct {
	i, j int
}

// 82/84
func bfs(nodes map[Node][]Node, start, finish Node) int {
	visited := make(map[Node]bool)

	current := []Node{start}
	path := 1
	for len(current) != 0 {
		path++
		for _, node := range current {
			if node == finish {
				return path
			}

			for _, friend := range nodes[node] {
				if !visited[friend] {
					visited[friend] = true
					current = append(current, friend)
				}
			}

			current = current[1:]
		}
	}

	return -1
}

func shortestPathBinaryMatrix(grid [][]int) int {
	nodes := make(map[Node][]Node)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				continue
			}

			node := Node{i, j} // loop : -1 .. 0 .. 1
			if j+1 < len(grid[i]) && grid[i][j+1] == 0 {
				nodes[node] = append(nodes[node], Node{i, j + 1})
			}

			if i+1 < len(grid) && grid[i+1][j] == 0 {
				nodes[node] = append(nodes[node], Node{i + 1, j})
			}

			if i+1 < len(grid) && j+1 < len(grid[i]) && grid[i+1][j+1] == 0 {
				nodes[node] = append(nodes[node], Node{i + 1, j + 1})
			}

			if i+1 < len(grid) && j-1 >= 0 && grid[i+1][j-1] == 0 {
				nodes[node] = append(nodes[node], Node{i + 1, j - 1})
			}

			if j-1 >= 0 && grid[i][j-1] == 0 {
				nodes[node] = append(nodes[node], Node{i, j - 1})
			}
		}
	}

	return bfs(nodes, Node{0, 0}, Node{len(grid) - 1, len(grid) - 1})
}

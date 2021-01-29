package jan

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val int
	Y   int
}

func verticalTraversal(root *TreeNode) [][]int {
	storage := make(map[int][]Node)
	min, max := walkHelper(root, 0, 0, storage)

	result := make([][]int, 0)
	for min <= max {
		if in, ok := storage[min]; ok {
			sort.Slice(in, func(i, j int) bool {
				if in[i].Y == in[j].Y {
					return in[i].Val < in[j].Val
				}

				return in[i].Y > in[j].Y
			})

			c := make([]int, len(in))
			for i := range c {
				c[i] = in[i].Val
			}

			result = append(result, c)
		}
		min++
	}

	return result
}

func walkHelper(root *TreeNode, x, y int, storage map[int][]Node) (int, int) {
	min, max := x, x
	storage[x] = append(storage[x], Node{Val: root.Val, Y: y})
	if root.Left != nil {
		lMin, lMax := walkHelper(root.Left, x-1, y-1, storage)

		if lMin < min {
			min = lMin
		}

		if lMax > max {
			max = lMax
		}
	}

	if root.Right != nil {
		rMin, rMax := walkHelper(root.Right, x+1, y-1, storage)

		if rMin < min {
			min = rMin
		}

		if rMax > max {
			max = rMax
		}
	}

	return min, max
}

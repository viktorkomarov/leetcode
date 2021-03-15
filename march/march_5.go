package march

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	result := make([]float64, 0)

	meta := make(map[int][]float64)
	walk(root, 0, meta)

	for i := 0; i < len(meta); i++ {
		result = append(result, average(meta[i]))
	}

	return result
}

func average(nums []float64) float64 {
	temp := 0.0
	for _, num := range nums {
		temp += num
	}

	return temp / float64(len(nums))
}

func walk(root *TreeNode, level int, iters map[int][]float64) {
	iters[level] = append(iters[level], float64(root.Val))

	if root.Left != nil {
		walk(root.Left, level+1, iters)
	}

	if root.Right != nil {
		walk(root.Right, level+1, iters)
	}
}

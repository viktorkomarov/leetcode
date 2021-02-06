package feb

func travel(node *TreeNode, depth int, rsv *[]int) {
	if node == nil {
		return
	}

	if depth > len(*rsv) {
		*rsv = append(*rsv, node.Val)
	}

	travel(node.Right, depth+1, rsv)
	travel(node.Left, depth+1, rsv)
}

func rightSideView(root *TreeNode) []int {
	rsv := make([]int, 0)
	depth := 0
	travel(root, depth+1, &rsv)
	return rsv
}

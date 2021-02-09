package feb

func walkGreater(root *TreeNode, acc int) int {
	if root == nil {
		return acc
	}

	right := walkGreater(root.Right, acc)
	root.Val += right
	if root.Left != nil {
		return walkGreater(root.Left, root.Val)
	}

	return root.Val
}

func convertBST(root *TreeNode) *TreeNode {
	walkGreater(root, 0)
	return root
}

package march

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d != 1 {
		rowAdd(root, v, d)
		return root
	}

	node := &TreeNode{Val: v}
	if v < root.Val {
		node.Right = root
	} else {
		node.Left = root
	}

	return node

}

func rowAdd(parent *TreeNode, v int, d int) {
	if d-1 == 1 {
		parent.Left = &TreeNode{Val: v, Left: parent.Left}
		parent.Right = &TreeNode{Val: v, Right: parent.Right}
		return
	}

	if parent.Left != nil {
		rowAdd(parent.Left, v, d-1)
	}

	if parent.Right != nil {
		rowAdd(parent.Right, v, d-1)
	}
}

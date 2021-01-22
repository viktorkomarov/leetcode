package leetcode

func deleteDuplicates(head *ListNode) *ListNode {
	skip := func(node *ListNode, val int) *ListNode {
		for node != nil && node.Val == val {
			node = node.Next
		}

		return node
	}

	result := ListNode{}
	current := &result

	for head != nil {
		next := head.Next
		if next == nil || next.Val != head.Val {
			current.Next = &ListNode{Val: head.Val}
			current = current.Next
		}

		head = skip(next, head.Val)
	}

	return result.Next
}

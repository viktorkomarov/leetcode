package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	result := ListNode{}
	filler := &result
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			filler.Val = l1.Val
			l1 = l1.Next
		} else {
			filler.Val = l2.Val
			l2 = l2.Next
		}

		filler.Next = &ListNode{}
		filler = filler.Next
	}

	copied := func(node *ListNode) {
		lazy := false
		for node != nil {
			if lazy {
				filler.Next = &ListNode{}
				filler = filler.Next
			}

			filler.Val = node.Val
			node = node.Next
			lazy = true
		}
	}

	copied(l1)
	copied(l2)

	return &result
}

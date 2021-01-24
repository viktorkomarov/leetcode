package leetcode

func canIterate(lists []*ListNode) int {
	result := -1
	for i := range lists {
		if lists[i] == nil {
			continue
		}

		if result == -1 || lists[result].Val > lists[i].Val {
			result = i
		}
	}

	return result
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head ListNode
	iterator := &head

	for {
		idx := canIterate(lists)
		if idx == -1 {
			return head.Next
		}

		iterator.Next = &ListNode{Val: lists[idx].Val}
		iterator = iterator.Next
		lists[idx] = lists[idx].Next
	}
}

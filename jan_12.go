package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := ListNode{}
	dummyIter := &dummy

	valAndStep := func(l *ListNode) (*ListNode, int) {
		if l == nil {
			return nil, 0
		}

		saved := l.Val
		l = l.Next
		return l, saved
	}

	var (
		overflow = 0
		val1     = 0
		val2     = 0
	)
	for l1 != nil || l2 != nil || overflow == 1 {
		l1, val1 = valAndStep(l1)
		l2, val2 = valAndStep(l2)

		newVal := val1 + val2 + overflow
		if newVal < 10 {
			overflow = 0
			dummyIter.Next = &ListNode{Val: newVal}
		} else {
			overflow = 1
			dummyIter.Next = &ListNode{Val: newVal - 10}
		}

		dummyIter = dummyIter.Next
	}

	return dummy.Next
}

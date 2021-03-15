package march

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	commons := make(map[*ListNode]bool)

	for headA != nil {
		commons[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if commons[headB] {
			return headB
		}

		headB = headB.Next
	}

	return nil
}

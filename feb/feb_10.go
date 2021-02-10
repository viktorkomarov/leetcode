package feb

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	save := head

	nodes := make(map[*Node]*Node)
	for head != nil {
		nodes[head] = &Node{
			Val:    head.Val,
			Next:   head.Next,
			Random: head.Random,
		}

		head = head.Next
	}

	dummy := Node{}
	iter := &dummy
	for save != nil {
		iter.Next = nodes[save]
		if save.Next != nil {
			iter.Next.Next = nodes[save.Next]
		}

		if save.Random != nil {
			iter.Next.Random = nodes[save.Random]
		}

		iter = iter.Next
		save = save.Next
	}

	return dummy.Next
}

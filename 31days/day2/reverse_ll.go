package day2

type Node struct {
	data int
	next *Node
}

func reverse_itr(head *Node) *Node {
	var previous, current *Node = nil, head
	for current != nil {
		tmp := current.next
		current.next = previous
		previous = current
		current = tmp
	}
	return previous
}

func reverse_rec(head *Node) *Node {
	if head == nil {
		return nil
	}
	head.next = reverse_rec(head.next)
	return head
}

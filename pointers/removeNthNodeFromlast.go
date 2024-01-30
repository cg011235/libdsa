package pointers

type LinkedListNode struct {
	data int
	next *LinkedListNode
}

func removeNthLastNode(head *LinkedListNode, n int) *LinkedListNode {
	if head == nil || n <= 0 {
		return head
	}
	dummy := &LinkedListNode{next: head}
	first, second := dummy, dummy
	for i := 0; i <= n; i++ {
		if first == nil {
			return head
		}
		first = first.next
	}
	for first != nil {
		first = first.next
		second = second.next
	}
	second.next = second.next.next
	return dummy.next
}

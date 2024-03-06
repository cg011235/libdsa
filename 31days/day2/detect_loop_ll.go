package day2

func detect_loop(head *Node) bool {
	if head == nil {
		return false
	}
	if head.next == nil {
		return false
	}
	slow, fast := head, head
	for {
		slow = slow.next
		if fast.next != nil && fast.next.next != nil {
			fast = fast.next.next
		} else {
			return false
		}
		if slow == fast {
			return true
		}
	}
}

package day2

func merge_ll(head1, head2 *Node) *Node {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	var newHead, current, previous *Node
	for head1 != nil && head2 != nil {
		if head1.data < head2.data {
			current = head1
			head1 = head1.next
		} else {
			current = head2
			head2 = head2.next
		}
		if newHead == nil {
			newHead = current
			previous = newHead
		} else {
			previous.next = current
			previous = current
		}
	}

	if head1 != nil {
		previous.next = head1
	}

	if head2 != nil {
		previous.next = head2
	}

	return newHead
}

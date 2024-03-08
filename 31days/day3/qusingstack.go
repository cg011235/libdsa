package day3

type stak struct {
	arr []int
}

func (s *stak) push(x int) {
	s.arr = append(s.arr, x)
}

func (s *stak) pop() (int, bool) {
	if len(s.arr) == 0 {
		return 0, false
	}
	elem := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return elem, true
}

func (s *stak) isEmpty() bool {
	return len(s.arr) == 0
}

type Q struct {
	st1 stak
	st2 stak
}

func (q *Q) enqueue(x int) {
	q.st1.push(x)
}

func (q *Q) dequeue() (int, bool) {
	if q.st2.isEmpty() {
		if q.st1.isEmpty() {
			// Both stacks are empty, indicating the queue is empty
			return 0, false
		}
		// Transfer all elements from st1 to st2, reversing their order
		for !q.st1.isEmpty() {
			elem, _ := q.st1.pop() // Safe to ignore the bool return value here due to isEmpty check
			q.st2.push(elem)
		}
	}
	return q.st2.pop() // Now st2 is guaranteed to have elements
}

func (q *Q) isEmpty() bool {
	return q.st1.isEmpty() && q.st2.isEmpty()
}

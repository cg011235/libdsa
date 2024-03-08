package day3

type MinStack struct {
	mainStack []int // Stack to store all elements
	minStack  []int // Stack to store minimum elements
}

// Push adds an element onto the stack
func (ms *MinStack) Push(x int) {
	ms.mainStack = append(ms.mainStack, x)
	// If minStack is empty or x is less or equal to the current minimum, push it onto minStack
	if len(ms.minStack) == 0 || x <= ms.GetMin() {
		ms.minStack = append(ms.minStack, x)
	}
}

// Pop removes the top element from the stack and returns it
func (ms *MinStack) Pop() (int, bool) {
	if len(ms.mainStack) == 0 {
		return 0, false
	}
	topElement := ms.mainStack[len(ms.mainStack)-1]
	ms.mainStack = ms.mainStack[:len(ms.mainStack)-1]

	// If the popped element is the current minimum, pop it from minStack as well
	if topElement == ms.minStack[len(ms.minStack)-1] {
		ms.minStack = ms.minStack[:len(ms.minStack)-1]
	}

	return topElement, true
}

// Top returns the top element of the stack
func (ms *MinStack) Top() (int, bool) {
	if len(ms.mainStack) == 0 {
		return 0, false
	}
	return ms.mainStack[len(ms.mainStack)-1], true
}

// GetMin returns the minimum element in the stack
func (ms *MinStack) GetMin() int {
	if len(ms.minStack) == 0 {
		return 0 // Assuming 0 as a default value. Adjust based on the use case or add error handling.
	}
	return ms.minStack[len(ms.minStack)-1]
}

// IsEmpty checks if the stack is empty
func (ms *MinStack) IsEmpty() bool {
	return len(ms.mainStack) == 0
}

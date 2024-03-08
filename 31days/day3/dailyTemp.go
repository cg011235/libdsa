package day3

type TemperatureStack struct {
	arr []int
}

func (ts *TemperatureStack) push(x int) {
	ts.arr = append(ts.arr, x)
}

func (ts *TemperatureStack) pop() (int, bool) {
	if len(ts.arr) == 0 {
		return -1, false
	}
	element := ts.arr[len(ts.arr)-1]
	ts.arr = ts.arr[:len(ts.arr)-1]
	return element, true
}

func (ts *TemperatureStack) isEmpty() bool {
	return len(ts.arr) == 0
}

func (ts *TemperatureStack) top() (int, bool) {
	if len(ts.arr) == 0 {
		return -1, false
	}
	return ts.arr[len(ts.arr)-1], true
}

func dailyTemperatures(T []int) []int {
	n := len(T)
	out := make([]int, n) // Initialize output array with zeros
	var stack TemperatureStack

	for i, temperature := range T {
		// Check if current temperature is warmer than the top of the stack
		for !stack.isEmpty() {
			prevIndex, _ := stack.top()
			if T[prevIndex] < temperature {
				// If warmer, update the output array and pop from the stack
				stack.pop()
				out[prevIndex] = i - prevIndex
			} else {
				// If not warmer, break from the loop
				break
			}
		}
		// Push the current index onto the stack
		stack.push(i)
	}

	return out
}

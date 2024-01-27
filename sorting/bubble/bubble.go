package bubble

func BubbleSort(arr []int) []int {
	n := len(arr)
	already_sorted := false
	for i := 0; i < (n-1) && already_sorted == false; i++ {
		already_sorted = true
		// no swapping of values for one round means
		// array is already sorted.
		for j := 0; j < (n - i - 1); j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
				already_sorted = false // swapped means not sorted yet.
			}
		}
	}
	return arr
}

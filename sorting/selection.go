package sorting

func selectionSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		max := i
		for j := i + 1; j < n-i; j++ {
			if arr[j] > arr[max] {
				max = j
			}
		}
		if max != n-i-1 {
			arr[n-i-1], arr[max] = arr[max], arr[n-i-1]
		}
	}
	return arr
}

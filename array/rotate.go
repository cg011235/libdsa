package array

func rotate(arr []int, k int) {
	n := len(arr)
	for i := 0; i < k; i++ {
		tmp := arr[0]
		for j := 0; j < n-1; j++ {
			arr[j] = arr[j+1]
		}
		arr[n-1] = tmp
	}
}

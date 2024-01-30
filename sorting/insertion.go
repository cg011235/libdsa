package sorting

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		// Move elements of arr[0..i-1],
		// that are greater than key,
		// to one position ahead
		// of their current position
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

/*
20, 10

key = 10
j = 0
20 > 10
arr[1] = 20
j = -1
arr[0] = key = 10

20, 10, 8, 30
*/

package bubble

import "testing"

func TestBasic(t *testing.T) {
	arr := []int{7, 5, 8, 3}
	result := BubbleSort(arr)
	n := len(result)
	for i := 0; i < n-1; i++ {
		if result[i] > result[i+1] {
			t.Error("Failed sort array")
		}
	}

	arr = []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	result = BubbleSort(arr)
	n = len(result)
	for i := 0; i < n-1; i++ {
		if result[i] > result[i+1] {
			t.Error("Failed sort array")
		}
	}
}

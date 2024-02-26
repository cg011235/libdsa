package day1

import (
	"reflect"
	"sort"
	"testing"
)

/*
  Given an array nums of n integers, are there elements a, b, c in nums
	such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero
*/

func sumOf3(arr []int) (triplets [][]int) {
	sort.Ints(arr)
	n := len(arr)

	for i := 0; i < n; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			if sum == 0 {
				triplets = append(triplets, []int{arr[i], arr[left], arr[right]})

				// Skip duplicates for the second element
				for left < right && arr[left] == arr[left+1] {
					left++
				}
				// Skip duplicates for the third element
				for left < right && arr[right] == arr[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				// If the sum is less than 0, move the left pointer to increase the sum
				left++
			} else {
				// If the sum is more than 0, move the right pointer to decrease the sum
				right--
			}
		}
	}

	return
}

func TestSumOf3(t *testing.T) {
	arr := []int{-4, -1, 0, 1, 2}
	triplets := sumOf3(arr)
	expected := [][]int{{-1, 0, 1}}

	if len(triplets) != len(expected) {
		t.Errorf("Expected %d triplets, got %d", len(expected), len(triplets))
	}

	for i, triplet := range triplets {
		if !reflect.DeepEqual(triplet, expected[i]) {
			t.Errorf("Expected triplet %v, got %v", expected[i], triplet)
		}
	}

	for _, triplet := range triplets {
		t.Logf("Triplet: %v", triplet)
	}
}

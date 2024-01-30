package pointers

import (
	"sort"
)

// if there are any three integers in nums
// whose sum is equal to the target
func findSumOfThree(nums []int, target int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		low, high := i+1, len(nums)-1
		for low < high {
			sum := nums[i] + nums[low] + nums[high]
			if sum == target {
				return true
			} else if sum < target {
				low++
			} else {
				high--
			}
		}
	}
	return false
}

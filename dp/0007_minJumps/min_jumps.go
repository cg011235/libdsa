package minjumps

import "math"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minJumpsBottomUp(arr []int) int {
	n := len(arr)
	if n == 0 || arr[0] == 0 {
		return math.MaxInt32
	}

	jumps := make([]int, n)
	for i := range jumps {
		jumps[i] = math.MaxInt32
	}
	jumps[0] = 0

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if i <= j+arr[j] && jumps[j] != math.MaxInt32 {
				jumps[i] = min(jumps[i], jumps[j]+1)
				break
			}
		}
	}

	if jumps[n-1] == math.MaxInt32 {
		return -1 // Not possible to reach the end
	}

	return jumps[n-1]
}

func minJumpsTopDown(arr []int) int {
	n := len(arr)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	return jumpHelper(0, arr, memo)
}

// jumpHelper is the recursive function to calculate the minimum jumps.
func jumpHelper(position int, arr []int, memo []int) int {
	// Base case: when at or beyond the last element.
	if position >= len(arr)-1 {
		return 0
	}

	// Return memoized result if already calculated.
	if memo[position] != -1 {
		return memo[position]
	}

	// If the current element is 0, it's not possible to move forward.
	if arr[position] == 0 {
		return math.MaxInt32
	}

	minJumps := math.MaxInt32
	for i := 1; i <= arr[position]; i++ {
		jumps := jumpHelper(position+i, arr, memo)
		if jumps != math.MaxInt32 {
			minJumps = min(minJumps, jumps+1)
		}
	}

	// Memoize the result before returning.
	memo[position] = minJumps
	return minJumps
}

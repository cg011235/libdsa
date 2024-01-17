package eggdropping

import (
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func eggDropTopDown(eggs, floors int, memo [][]int) int {
	if floors == 0 || floors == 1 {
		return floors
	}

	if eggs == 1 {
		return floors
	}

	if memo[eggs][floors] != -1 {
		return memo[eggs][floors]
	}

	min := math.MaxInt32

	for k := 1; k <= floors; k++ {
		res := max(eggDropTopDown(eggs-1, k-1, memo), eggDropTopDown(eggs, floors-k, memo))
		if res < min {
			min = res
		}
	}

	memo[eggs][floors] = min + 1
	return memo[eggs][floors]
}

func eggDropBottomUp(eggs, floors int) int {
	dp := make([][]int, eggs+1)
	for i := range dp {
		dp[i] = make([]int, floors+1)
	}

	// We need one trial for one floor and zero trials for zero floors
	for i := 1; i <= eggs; i++ {
		dp[i][0] = 0
		dp[i][1] = 1
	}

	// We always need j trials for one egg and j floors.
	for j := 1; j <= floors; j++ {
		dp[1][j] = j
	}

	for i := 2; i <= eggs; i++ {
		for j := 2; j <= floors; j++ {
			dp[i][j] = math.MaxInt32
			for k := 1; k <= j; k++ {
				res := 1 + max(dp[i-1][k-1], dp[i][j-k])
				if res < dp[i][j] {
					dp[i][j] = res
				}
			}
		}
	}

	return dp[eggs][floors]
}

package knapsack

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func knapsack(dp [][]int, W, V []int, n, c int) int {
	// base condition
	if n == 0 || c == 0 {
		return 0
	}

	if dp[n][c] != -1 {
		return dp[n][c]
	}

	if W[n-1] > c {
		// Item excluded from knapsack because
		// its weight it greater than knapsack
		// capacity.
		return knapsack(dp, W, V, n-1, c)
	} else {
		// Possiblity of item skipped from knapsack.
		tmp1 := knapsack(dp, W, V, n-1, c)
		// Possiblity of item kept in knapsack.
		tmp2 := V[n-1] + knapsack(dp, W, V, n-1, c-W[n-1])
		dp[n][c] = max(tmp1, tmp2)
		return dp[n][c]
	}
}

func knapsack_dp(W, V []int, n, c int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, c+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return knapsack(dp, W, V, n, c)
}

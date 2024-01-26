package lcs

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lcs(dp [][]int, p, q string, n, m int) int {
	if n == 0 || m == 0 {
		return 0
	}

	if dp[n][m] != -1 {
		return dp[n][m]
	}

	var result int
	if p[n-1] == q[m-1] {
		result = 1 + lcs(dp, p, q, n-1, m-1)
	} else {
		tmp1 := lcs(dp, p, q, n-1, m)
		tmp2 := lcs(dp, p, q, n, m-1)
		result = max(tmp1, tmp2)
	}
	dp[n][m] = result
	return result
}

func lcs_dp(p, q string) int {
	n := len(p)
	m := len(q)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return lcs(dp, p, q, n, m)
}

func lcs_bottom_up(p, q string) int {
	n := len(p)
	m := len(q)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if p[i-1] == q[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}

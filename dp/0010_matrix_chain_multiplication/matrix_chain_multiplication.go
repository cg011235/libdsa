package matrixchainmultiplication

import (
	"fmt"
	"math"
)

type Matrix struct {
	rows, cols int
}

func minMatrixChainOrder(matrices []Matrix) (int, string) {
	n := len(matrices)
	memo := make([][]int, n)
	split := make([][]int, n)

	for i := range memo {
		memo[i] = make([]int, n)
		split[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	cost := matrixChainOrderHelper(matrices, 0, n-1, memo, split)
	order := constructOptimalOrder(split, 0, n-1)
	return cost, order
}

func matrixChainOrderHelper(matrices []Matrix, i, j int, memo, split [][]int) int {
	if i == j {
		return 0
	}

	if memo[i][j] != -1 {
		return memo[i][j]
	}

	memo[i][j] = math.MaxInt32

	for k := i; k < j; k++ {
		count := matrixChainOrderHelper(matrices, i, k, memo, split) +
			matrixChainOrderHelper(matrices, k+1, j, memo, split) +
			matrices[i].rows*matrices[k].cols*matrices[j].cols

		if count < memo[i][j] {
			memo[i][j] = count
			split[i][j] = k
		}
	}

	return memo[i][j]
}

func constructOptimalOrder(split [][]int, i, j int) string {
	if i == j {
		return fmt.Sprintf("A%d", i+1)
	}
	return "(" + constructOptimalOrder(split, i, split[i][j]) + " x " +
		constructOptimalOrder(split, split[i][j]+1, j) + ")"
}

func minMatrixChainOrderBU(matrices []Matrix) int {
	n := len(matrices)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Cost is zero when multiplying one matrix
	for i := 0; i < n; i++ {
		dp[i][i] = 0
	}

	// L is the chain length
	for L := 2; L <= n; L++ {
		for i := 0; i < n-L+1; i++ {
			j := i + L - 1
			dp[i][j] = math.MaxInt32
			for k := i; k < j; k++ {
				cost := dp[i][k] + dp[k+1][j] + matrices[i].rows*matrices[k].cols*matrices[j].cols
				if cost < dp[i][j] {
					dp[i][j] = cost
				}
			}
		}
	}

	return dp[0][n-1]
}

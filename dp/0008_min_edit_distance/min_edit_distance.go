package mineditdistance

import "fmt"

func minof(x, y, z int) int {
	if x < y {
		if x < z {
			return x
		} else {
			return z
		}
	} else {
		if y < z {
			return y
		} else {
			return z
		}
	}
}

func minEditDistanceTopDown(str1, str2 string) int {
	memo := make(map[string]int)
	return editDistanceHelper(str1, str2, memo)
}

func editDistanceHelper(str1, str2 string, memo map[string]int) int {
	// Create a unique key for memoization
	key := fmt.Sprintf("%s-%s", str1, str2)

	if val, exists := memo[key]; exists {
		return val
	}

	if len(str1) == 0 {
		return len(str2)
	}
	if len(str2) == 0 {
		return len(str1)
	}

	// If last characters are same, ignore them and recurse for remaining strings
	if str1[len(str1)-1] == str2[len(str2)-1] {
		memo[key] = editDistanceHelper(str1[:len(str1)-1], str2[:len(str2)-1], memo)
	} else {
		// If last characters are not the same, consider all three operations and choose the one with minimum cost
		insertOp := editDistanceHelper(str1, str2[:len(str2)-1], memo)
		deleteOp := editDistanceHelper(str1[:len(str1)-1], str2, memo)
		replaceOp := editDistanceHelper(str1[:len(str1)-1], str2[:len(str2)-1], memo)

		memo[key] = 1 + minof(insertOp, deleteOp, replaceOp)
	}

	return memo[key]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minEditDistanceBottomUp(str1, str2 string) int {
	lenStr1, lenStr2 := len(str1), len(str2)
	dp := make([][]int, lenStr1+1)

	for i := range dp {
		dp[i] = make([]int, lenStr2+1)
	}

	// Fill the first row and first column of the dp table
	for i := 0; i <= lenStr1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= lenStr2; j++ {
		dp[0][j] = j
	}

	// Build the dp table
	for i := 1; i <= lenStr1; i++ {
		for j := 1; j <= lenStr2; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] // Characters are the same, no operation needed
			} else {
				dp[i][j] = min(
					dp[i-1][j]+1, // Deletion
					min(
						dp[i][j-1]+1,   // Insertion
						dp[i-1][j-1]+1, // Substitution
					),
				)
			}
		}
	}

	return dp[lenStr1][lenStr2]
}

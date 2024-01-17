package coinchanging

import "math"

func minCoinsTopDown(coins []int, amount int) int {
	memo := make(map[int]int)
	return coinChangeHelper(coins, amount, memo)
}

func coinChangeHelper(coins []int, amount int, memo map[int]int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if val, exists := memo[amount]; exists {
		return val
	}

	minCoins := math.MaxInt32
	for _, coin := range coins {
		res := coinChangeHelper(coins, amount-coin, memo)
		if res >= 0 && res < minCoins {
			minCoins = res + 1
		}
	}

	if minCoins == math.MaxInt32 {
		memo[amount] = -1
	} else {
		memo[amount] = minCoins
	}
	return memo[amount]
}

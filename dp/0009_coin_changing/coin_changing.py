#!/usr/bin/python

import sys
import unittest

def coin_change(coins, amount, memo=None):
    if memo is None:
        memo = {}
    if amount < 0:
        return -1
    if amount == 0:
        return 0
    if amount in memo:
        return memo[amount]
    min_coins = float('inf')
    for coin in coins:
        if coin > amount:
            continue
        num_coins = coin_change(coins, amount - coin, memo)  # current coint is used
        if 0 <= num_coins < min_coins:
            min_coins = 1 + num_coins
    memo[amount] = -1 if min_coins == float('inf') else min_coins
    return memo[amount]

class TestCoinChange(unittest.TestCase):

    def test_simple_case(self):
        self.assertEqual(coin_change([1, 2, 5], 11), 3)  # 5 + 5 + 1

    def test_no_coins(self):
        self.assertEqual(coin_change([], 11), -1)

    def test_zero_amount(self):
        self.assertEqual(coin_change([1, 2, 5], 0), 0)

    def test_no_solution(self):
        self.assertEqual(coin_change([2], 3), -1)

    def test_single_coin(self):
        self.assertEqual(coin_change([1], 100), 100)

    def test_large_amount(self):
        self.assertEqual(coin_change([1, 5, 10, 25], 100), 4)  # 25 * 4

if __name__ == '__main__':
    unittest.main()


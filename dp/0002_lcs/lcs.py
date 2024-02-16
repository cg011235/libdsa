#!/usr/bin/python

import unittest

def max(a, b):
    return a if a >= b else b

def lcs(P, Q, m, n, memo):
    if m == 0 or n == 0:
        return 0

    if memo[m-1][n-1] != -1:
        return memo[m-1][n-1]

    if P[m-1] == Q[n-1]:
        memo[m-1][n-1] = 1 + lcs(P, Q, m-1, n-1, memo)
    else:
        tmp1 = lcs(P, Q, m-1, n, memo)
        tmp2 = lcs(P, Q, m, n-1, memo)
        memo[m-1][n-1] = max(tmp1, tmp2)

    return memo[m-1][n-1]

def lcs_main(P, Q):
    m, n = len(P), len(Q)
    memo = [[-1 for _ in range(n)] for _ in range(m)]
    return lcs(P, Q, m, n, memo)

class TestLCS(unittest.TestCase):
    def test_lcs(self):
        self.assertEqual(lcs_main("ABCBDAB", "BDCABC"), 4)  # LCS is "BCAB" or "BDAB"
        self.assertEqual(lcs_main("", ""), 0)
        self.assertEqual(lcs_main("AXYT", "AYZX"), 2)  # LCS is "AY"
        self.assertEqual(lcs_main("12345", "54321"), 1)  # LCS could be any single digit
        self.assertEqual(lcs_main("AAAA", "AAAA"), 4)  # LCS is "AAAA"

if __name__ == "__main__":
    unittest.main()


#include <algorithm>
#include <string>
#include <vector>

using namespace std;

int lcs(vector<vector<int>>& dp, const string& p, const string& q, int m, int n) {
  if (m == 0 || n == 0) {
    return 0;
  }
  if (dp[m - 1][n - 1] != -1) {
    return dp[m - 1][n - 1];
  }
  if (p[m - 1] == q[n - 1]) {
    dp[m - 1][n - 1] = 1 + lcs(dp, p, q, m - 1, n -1);
    return dp[m - 1][n - 1];
  }
  int tmp1 = lcs(dp, p, q, m - 1, n);
  int tmp2 = lcs(dp, p, q, m, n - 1);
  dp[m - 1][n - 1] = max(tmp1, tmp2);
  return dp[m - 1][n - 1];
}


int lcs_main(const string& p, const string& q) {
    int m = p.size();
    int n = q.size();
    vector<vector<int>> dp(m + 1, vector<int>(n + 1, -1)); // Initialize all elements to -1
    return lcs(dp, p, q, m, n);
}
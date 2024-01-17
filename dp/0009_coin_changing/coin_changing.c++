#include <limits>
#include <map>

using namespace std;

map<int, int> memo;

int coinChange(int coins[], int n, int amount) {
  if (amount < 0) {
    return -1;
  }
  if (amount == 0) {
    return 0;
  }
  if (n <= 0) {
    return 0;
  }

  if (memo.find(amount) != memo.end()) {
    return memo[amount];
  }

  int minCoins = __INT_MAX__;
  for (int i = 0; i < n; i++) {
    if (coins[i] > amount) {
      continue;
    }
    int result = coinChange(coins, n, amount - coins[i]);
    if (result >= 0 && result < minCoins) {
      minCoins = result + 1;
    }
  }

  if (minCoins == __INT_MAX__) {
    memo[amount] = -1;
  } else {
    memo[amount] = minCoins;
  }
  return memo[amount];
}
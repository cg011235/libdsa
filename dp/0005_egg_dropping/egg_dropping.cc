#include <climits>
#include <vector>

using namespace std;

int egg_dropping(int eggs, int floors, vector<vector<int>> memo) {
  if (floors == 0 || floors == 1) {
    return floors;
  }
  if (eggs == 1) {
    return floors;
  }
  if (memo[eggs][floors] != -1) {
    return memo[eggs][floors];
  }
  int min = INT_MAX;
  for (int i = 0; i < floors; i++) {
    int tmp1 = egg_dropping(eggs - 1, i - 1, memo);
    int tmp2 = egg_dropping(eggs, floors - i, memo);
    int res = max(tmp1, tmp2);
    if (res < min) {
      min = res;
    }
  }
  memo[eggs][floors] = min + 1;
  return memo[eggs][floors];
}
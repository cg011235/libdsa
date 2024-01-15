#include "01knapsack.h++"

#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

// Recursive function with memoization
int knapsackRecursive(const vector<int>& values, const vector<int>& weights, int capacity, int index, vector<vector<int>>& memo) {
    if (index == 0 || capacity == 0) {
        return 0;
    }

    // Check if value is already computed
    if (memo[index][capacity] != -1) {
        return memo[index][capacity];
    }

    // If weight of the current item is more than the capacity, skip it
    if (weights[index - 1] > capacity) {
        memo[index][capacity] = knapsackRecursive(values, weights, capacity, index - 1, memo);
        return memo[index][capacity];
    }

    // Possiblities of inclusion of the current item and exclusion of the current item
    int include = values[index - 1] + knapsackRecursive(values, weights, capacity - weights[index - 1], index - 1, memo);
    int exclude = knapsackRecursive(values, weights, capacity, index - 1, memo);

    memo[index][capacity] = max(include, exclude);
    return memo[index][capacity];
}

// Interface to expose.
int knapsack(const vector<int>& values, const vector<int>& weights, int capacity) {
    int n = values.size();
    vector<vector<int>> memo(n + 1, vector<int>(capacity + 1, -1));
    return knapsackRecursive(values, weights, capacity, n, memo);
}

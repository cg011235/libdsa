#include <vector>
#include <algorithm> // For std::rotate

void rotate(std::vector<int>& arr, int k) {
    int n = arr.size(); // Use size() method to get the number of elements in the vector
    k = k % n; // In case k is greater than n, this will ensure we do the minimum necessary rotations

    // Use std::rotate from the algorithm header to rotate the vector
    // This rotates the range [first, n-first) to the beginning of the vector
    std::rotate(arr.begin(), arr.begin() + k, arr.end());
}

void rotate(int arr[], int k) {
  int n = sizeof(arr)/sizeof(arr[0]);
  k = k % n;
  for (int i = 0; i < k; i++) {
    int tmp = arr[0];
    for (int j = 0; j < n - 1; j++) {
      arr[j] = arr[j + 1];
    }
    arr[n - 1] = tmp;
  }
}

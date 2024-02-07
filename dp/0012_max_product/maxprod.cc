#include <vector>
#include <algorithm> // For std::max and std::min

int maxProduct(std::vector<int>& nums) {
    if (nums.empty()) return 0;

    int maxProd = nums[0], minProd = nums[0], result = nums[0];
    for (int i = 1; i < nums.size(); ++i) {
        if (nums[i] < 0) std::swap(maxProd, minProd);
        maxProd = std::max(nums[i], nums[i] * maxProd);
        minProd = std::min(nums[i], nums[i] * minProd);
        result = std::max(result, maxProd);
    }
    return result;
}

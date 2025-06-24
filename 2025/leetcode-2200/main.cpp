class Solution {
public:
  vector<int> findKDistantIndices(vector<int> &nums, int key, int k) {
    std::vector<int> res;
    int n = nums.size();
    for (int i = 0; i < n; i++) {
      if (nums[i] == key) {
        for (int j = std::max(0, i - k); j <= std::min(n - 1, i + k); j++) {
          if (res.empty() || j > res.back())
            res.emplace_back(j);
        }
      }
    }
    return res;
  }
};

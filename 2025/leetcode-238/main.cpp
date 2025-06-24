class Solution {
public:
  vector<int> productExceptSelf(vector<int> &nums) {
    int n = nums.size();
    std::vector<int> res(n);
    std::vector<int> prefix(n);
    std::vector<int> suffix(n);
    prefix[0] = nums[0];
    for (int i = 1; i < n; i++) {
      prefix[i] = nums[i] * prefix[i - 1];
    }
    suffix.back() = nums.back();
    for (int i = n - 2; i >= 0; i--) {
      suffix[i] = suffix[i + 1] * nums[i];
    }
    for (int i = 1; i < n - 1; i++) {
      res[i] = (prefix[i - 1] * suffix[i + 1]);
    }
    res[0] = suffix[1];
    res.back() = prefix[n - 2];
    return res;
  }
};

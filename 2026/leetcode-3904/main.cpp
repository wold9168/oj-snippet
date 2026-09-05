#include <algorithm>
#include <bits/stdc++.h>
#include <climits>
#include <vector>
using namespace std;
class Solution {
private:
  vector<int> max_to_left;  // max_to_left[i] means max(nums[0..i])
  vector<int> min_to_right; // min_to_right[i] means min(nums[i..n - 1])
  int n;
  void init(const vector<int> &nums) {
    n = nums.size();
    max_to_left = vector<int>(n);
    min_to_right = vector<int>(n);
    int recorded_max = INT_MIN;
    for (int i = 0; i < n; i++) {
      recorded_max = max(recorded_max, nums[i]);
      max_to_left[i] = recorded_max;
      cout << "[max] " << recorded_max << std::endl;
    }
    int recorded_min = INT_MAX;
    for (int i = n - 1; i >= 0; i--) {
      recorded_min = min(recorded_min, nums[i]);
      min_to_right[i] = recorded_min;
      cout << "[min] " << recorded_min << std::endl;
    }
  }
  int unstable_value(int i) { return max_to_left[i] - min_to_right[i]; }

public:
  int firstStableIndex(vector<int> &nums, int k) {
    init(nums);
    int res = -1;
    for (int i = 0; i < n; i++) {
      if (unstable_value(i) <= k) {
        res = i;
        break;
      }
      // cout<<i<<':'<<max_to_left[i]<<' ' << min_to_right[i]<<std::endl;
    }
    return res;
  }
};
int main() {
  Solution s;
  vector<int> test1 = {5, 0, 1, 3};
  s.firstStableIndex(test1, 3);
}

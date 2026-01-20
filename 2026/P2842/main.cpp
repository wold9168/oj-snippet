#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;
const int MAXN = 1e4 + 5;
int dp[MAXN] = {};
int main() {
  int n, m;
  cin >> n >> m;
  std::vector<int> vibi;
  for (int i = 0; i < n; i++) {
    int tmp;
    cin >> tmp;
    vibi.push_back(tmp);
  }
  sort(vibi.begin(), vibi.end());
  dp[0] = 0;
  for (int i = 1; i <= m; i++) {
    bool isValid = false;
    vector<int> restIndexList;
    for (int j = 0; j < vibi.size(); j++) {
      const auto to = vibi[j];
      auto rest = i - to;
      if (rest < 0 || dp[rest] == -1) {
        continue;
      } else {
        isValid = true;
        restIndexList.push_back(rest);
      }
    }
    if (isValid) {
      int minv = 0x1f1f1f1f;
      int mini = 0;
      for (auto ix : restIndexList) {
        if (dp[ix] < minv) {
          mini = ix;
          minv = dp[ix];
        }
      }
      dp[i] = minv + 1;
    } else
      dp[i] = -1;
  }
  cout << dp[m];
}

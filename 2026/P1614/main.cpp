#include <algorithm>
#include <array>
#include <iostream>

using namespace std;

int main() {
  int n, m;
  cin >> n >> m;
  array<int, 4096> arr;
  for (int i = 0; i < n; i++) {
    cin >> arr[i];
  }
  int imin = 0;
  int sum = 0;
  for (int i = 0; i < n - (m - 1); i++) {
    int L = i;
    int R = i + m - 1;
    if (i == 0) {
      for (int j = L; j <= R; j++) {
        sum += arr[j];
      }
      imin = sum;
    } else {
      sum = sum - arr[L - 1] + arr[R];
      imin = min(sum, imin);
    }
  }
  cout << imin << '\n';
  return 0;
}

#include <iostream>
#include <vector>
using namespace std;

vector<vector<int>> m;
int Mx = 0, My = 0;
void place(int x) {
  m[Mx][My] = x;
  My++;
  if (My == m.size()) {
    My = 0, Mx++;
  }
}
void insert(bool zero, int x) {
  int p = zero ? 0 : 1;
  for (int i = 0; i < x; i++) {
    place(p);
  }
  return;
}
int main() {
  int N;
  cin >> N;
  m.resize(N);
  for (auto &to : m) {
    to.resize(N);
  }
  int x;
  bool zero = true;
  while (cin >> x) {
    insert(zero, x);
    zero = !zero;
  }
  for (auto &to : m) {
    for (auto &x : to) {
      cout << x;
    }
    cout << '\n';
  }
}

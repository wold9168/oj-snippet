#include <iostream>
#include <string>
#include <vector>
using namespace std;
int main() {
  string lines;
  std::getline(std::cin, lines);
  int N = lines.size();
  for (int i = 1; i < N; i++) {
    string line;
    std::getline(std::cin, line);
    lines += line;
  }
  bool iszero = true;
  int x = 0;
  vector<int> A;
  for (auto to : lines) {
    if (to == '0') {
      if (iszero == true)
        x++;
      else {
        A.push_back(x);
        x = 1;
        iszero = !iszero;
      }
    } else if (to == '1') {
      if (iszero == false) {
        x++;
      } else {
        A.push_back(x);
        x = 1;
        iszero = !iszero;
      }
    } else {
      throw "123";
    }
  }
  A.push_back(x);
  std::cout << N << " ";
  for (auto to : A) {
    std::cout << to << " ";
  }
}

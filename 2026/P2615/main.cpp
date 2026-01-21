#include <iostream>
#include <utility>
#include <vector>
template <class T> void vecprint(const std::vector<T> &vec) {
  for (const auto &to : vec) {
    std::cout << to << ' ';
  }
  std::cout << '\n';
}
template <class T> void vec2print(const std::vector<T> &vec) {
  for (const auto &to : vec) {
    vecprint(to);
  }
}

auto convert(std::pair<int, int> raw) -> decltype(auto) {
  return std::pair<int, int>(raw.first - 1, raw.second - 1);
}

void set(const std::pair<int, int> pos, std::vector<std::vector<int>> &map,
         int value) {
  map[pos.first][pos.second] = value;
  return;
}
int get(const std::pair<int, int> pos, std::vector<std::vector<int>> &map) {
  return map[pos.first][pos.second];
}

// auto rule(const std::vector<std::vector<int>> &map,
//           const std::pair<int, int> &prevpos) -> decltype(auto) {
//   std::pair<int, int> curpos;
//   if (prevpos.first == 1 && prevpos.second != N) {
//     curpos.first = N;
//     curpos.second = prevpos.second + 1;
//   } else if (prevpos.second == N && prevpos.first != 1) {
//     curpos.second = 1;
//     curpos.first = prevpos.second - 1;
//   } else if (prevpos == std::pair{1, N}) {
//     curpos.first = prevpos.first + 1;
//     curpos.second = prevpos.second;
//   } else {
//     if (map[prevpos.first - 1][prevpos.second + 1] == 0) {
//       curpos.first = prevpos.first - 1;
//       curpos.second = prevpos.second + 1;
//     } else {
//       curpos.first = prevpos.first + 1;
//       curpos.second = prevpos.second;
//     }
//   }
//   return curpos;
// }

int main() {
  int N;
  std::cin >> N;
  std::vector<std::vector<int>> M(N, std::vector<int>(N));
  if (N % 2 != 1) {
    throw "invalid input";
  }
  auto rule = [&](const std::pair<int, int> &prevpos) -> std::pair<int, int> {
    std::pair<int, int> respos;
    if (prevpos.first == 1 && prevpos.second != N) {
      respos = {N, prevpos.second + 1};
    } else if (prevpos.second == N && prevpos.first != 1) {
      respos = {prevpos.first - 1, 1};
    } else if (prevpos == std::pair{1, N}) {
      respos = {prevpos.first + 1, prevpos.second};
    } else {
      if (get(convert(std::pair{prevpos.first - 1, prevpos.second + 1}), M) ==
          0) {
        respos = {prevpos.first - 1, prevpos.second + 1};
      } else {
        respos = {prevpos.first + 1, prevpos.second};
      }
    }
    return respos;
  };
  std::pair<int, int> first = {1, (1 + N) / 2};
  set(convert(first), M, 1);
  std::pair<int, int> prevpos = first;
  for (int i = 2; i <= N * N; i++) {
    std::pair<int, int> curpos = rule(prevpos);
    set(convert(curpos), M, i);
    prevpos = curpos;
  }
  vec2print(M);
}

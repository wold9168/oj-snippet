#include <bits/stdc++.h>
using namespace std;
class Solution {
public:
  inline int credit(const char &x) {
    return x == 'X' ? 1 : (x == 'Y' ? -1 : 0);
  }
  int numberOfSubmatrices(vector<vector<char>> &grid) {
    vector<vector<int>> mpi = vector<vector<int>>();
    int res = 0;
    for (const auto &it : grid) {
      vector<int> tmp;
      for (const auto &it2 : it) {
        tmp.push_back(credit(it2));
      }
      mpi.push_back(tmp);
    }

    vector<vector<bool>> mpflag = vector<vector<bool>>();
    for (const auto &it : grid) {
      vector<bool> tmp;
      for (const auto &it2 : it) {
        tmp.push_back(it2=='X'?true:false);
      }
      mpflag.push_back(tmp);
    }

    for (int i = 1; i < mpi.size(); i++) {
      mpi[i][0] += mpi[i - 1][0];
      mpflag[i][0] = mpflag[i][0] || mpflag[i-1][0];
    }
    for(int j=1;j<mpi[0].size();j++)
    {mpi[0][j]+=mpi[0][j-1];
    mpflag[0][j] = mpflag[0][j] || mpflag[0][j-1];}
    for (int i = 1; i < mpi.size(); i++) {
      for (int j = 1; j < mpi[0].size(); j++) {
        mpi[i][j] += mpi[i-1][j]+ mpi[i][j - 1]-mpi[i-1][j-1];
        mpflag[i][j]=mpflag[i][j]||mpflag[i-1][j]||mpflag[i][j-1];
      }
    }

    for (int i = 0; i < mpi.size(); i++) {
      for (int j = 0; j < mpi[0].size(); j++) {
        if (mpi[i][j] == 0&&mpflag[i][j]==true) {
          res++;
        }
      }
    }
    return res;
  }
};
int main() {
  Solution s;
  vector<vector<char>> dat{{'X', 'Y', '.'}, {'Y', '.', '.'}};
  s.numberOfSubmatrices(dat);
}

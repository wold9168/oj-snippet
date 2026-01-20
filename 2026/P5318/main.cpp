#include <iostream>
#include <queue>
#include <set>
#include <stack>
using namespace std;
const int MAXN = 1e5 + 5;
set<int> graph[MAXN];

vector<int> bfstrace;
void bfs(int root = 1) {
  bool bfsnode_visited[MAXN] = {};
  queue<int> qu;
  qu.push(root);
  bfsnode_visited[root] = true;
  while (!qu.empty()) {
    auto curnode = qu.front();
    for (const auto &to : graph[curnode]) {
      if (bfsnode_visited[to] == false) {
        bfsnode_visited[to] = true;
        qu.push(to);
      }
    }
    bfstrace.push_back(curnode);
    qu.pop();
  }
  return;
}

vector<int> dfstrace_stack;
void dfs_stack(int root = 1) {
  bool dfsnode_visited[MAXN] = {};
  stack<int> st;
  st.push(root);
  // dfsnode_visited[root] = true;
  while (!st.empty()) {
    auto curnode = st.top();
    st.pop();
    if (dfsnode_visited[curnode] == true) {
      continue;
    }
    dfsnode_visited[curnode] = true;
    dfstrace_stack.push_back(curnode);
    for (auto to = graph[curnode].rbegin(); to != graph[curnode].rend(); ++to) {
      st.push(*to);
    }
  }
  return;
}

vector<int> dfstrace;
bool vis[MAXN] = {};
void dfs(const int root = 1) {
  if (vis[root])
    return;
  vis[root] = true;
  dfstrace.push_back(root);
  for (const auto &to : graph[root]) {
    dfs(to);
  }
}

void vecprint(const std::vector<int> &x) {
  for (const auto &to : x) {
    std::cout << to << ' ';
  }
  std::cout << '\n';
}

int main() {
  int n, m;
  cin >> n >> m;
  for (int i = 0; i < m; i++) {
    int a, b;
    cin >> a >> b;
    graph[a].insert(b);
  }
  dfs();
  dfs_stack();
  bfs();
  // vecprint(dfstrace);
  vecprint(dfstrace_stack);
  vecprint(bfstrace);
  return 0;
}

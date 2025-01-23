class Solution {
public:
  int countServers(vector<vector<int>> &grid) {
    vector<int> seen;
    auto n = grid.size();
    auto m = grid[0].size();

    seen.resize(m);
    auto result = 0;

    for (size_t i = 0; i < n; i++) {
      auto count = 0;
      auto only_one_pos = -1;

      for (size_t j = 0; j < m; j++) {
        if (grid[i][j] == 1) {
          count++;
          only_one_pos = j;
        }
      }

      if (count > 1) {
        result += count;
      }
      if (count == 1) {
        for (size_t k = 0; k < n; k++) {
          if (k == i)
            continue;
          if (grid[k][only_one_pos] == 1) {
            result++;
            break;
          }
        }
      }
    }

    return result;
  }
};
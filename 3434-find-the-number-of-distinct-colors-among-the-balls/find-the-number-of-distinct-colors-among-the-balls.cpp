class Solution {
public:
    vector<int> queryResults(int limit, vector<vector<int>>& queries) {
        vector<int> result;
        unordered_map<int, int> color;
        unordered_map<int, int> color_count;

        for (auto pair : queries) {
            auto p = pair[0];
            auto c = pair[1];

            if (color.count(p) > 0) {
                color_count[color[p]]--;
                if (color_count[color[p]] == 0)
                    color_count.erase(color[p]);
            }
            color_count[c]++;
            color[p] = c;

            result.push_back(color_count.size());
        }

        return result;
    }
};
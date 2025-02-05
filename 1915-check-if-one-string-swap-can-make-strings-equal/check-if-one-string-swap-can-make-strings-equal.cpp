class Solution {
public:
    bool areAlmostEqual(string s1, string s2) {
        auto c1 = 0;
        auto c2 = 0;
        auto diff = 0;
        for (auto i = 0; i < s1.size(); i++) {
            if (s1[i] != s2[i]) {
                diff++;
                if (diff > 2) {
                    return false;
                }
                if (diff == 1) {
                    c1 = i;
                }
                if (diff == 2) {
                    c2 = i;
                }
            }
        }
        return s1[c2] == s2[c1] && s1[c1] == s2[c2];
    }
};
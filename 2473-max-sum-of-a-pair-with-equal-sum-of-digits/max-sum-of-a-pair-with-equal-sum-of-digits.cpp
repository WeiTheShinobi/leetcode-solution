class Solution {
public:
    int maximumSum(vector<int>& nums) {
        auto result = -1;
        unordered_map<int, int> m;

        for (auto n : nums) {
            auto s = cal(n);
            if (m.count(s) > 0) {
                result = max(result, m[s] + n);
            }
            m[s] = max(m[s], n);
        }
        return result;
    }
    
    int cal(int n) {
        auto sum = 0;
        while (n > 0) {
            sum += n%10;
            n /= 10;
        }
        return sum;
    }
};
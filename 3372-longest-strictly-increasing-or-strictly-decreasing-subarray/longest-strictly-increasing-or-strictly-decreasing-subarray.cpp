class Solution {
public:
    int longestMonotonicSubarray(vector<int>& nums) {
        auto max_arr = 1;
        auto min_arr = 1;
        auto result = 1;
        for (auto i = 1; i < nums.size(); i++) {
            if (nums[i] > nums[i-1]) {
                max_arr++;
            } else {
                max_arr = 1;
            }
            if (nums[i] < nums[i-1]) {
                min_arr++;
            } else {
                min_arr = 1;
            }
            if (max_arr > result) {
                result = max_arr;
            }
            if (min_arr > result) {
                result = min_arr;
            }
        }

        return result;
    }
};
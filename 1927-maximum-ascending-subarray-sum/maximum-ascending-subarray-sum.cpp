class Solution {
public:
    int maxAscendingSum(vector<int>& nums) {
        auto result = nums[0];
        auto tmp = nums[0];
        for (auto i = 1; i < nums.size(); i++) {
            if (nums[i] > nums[i-1]) {
                tmp += nums[i];
            } else {
                tmp = nums[i];
            }
            result = max(result, tmp);
        }
        return result;      
    }
};
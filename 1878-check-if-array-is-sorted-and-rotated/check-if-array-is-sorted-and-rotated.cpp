class Solution {
public:
    bool check(vector<int>& nums) { 
        auto min_idx = 0;
        for (auto i = 0; i < nums.size(); i++) {
            if (nums[i] <= nums[min_idx]) {
                min_idx = i;
            }
            while (i < nums.size()-1 && nums[i] == nums[i+1]) i++;
        }

        for (auto i = 0; i < nums.size()-1; i++) {     
            auto next = (min_idx + 1) % nums.size();
            if (nums[next] < nums[min_idx]) {
                return false;
            }
            min_idx = next;
        }

        return true;
    }
};
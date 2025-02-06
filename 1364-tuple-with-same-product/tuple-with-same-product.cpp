class Solution {
public:
    int tupleSameProduct(vector<int>& nums) {
        auto result = 0;
        sort(nums.begin(), nums.end());
        
        for (auto i = 0; i < nums.size(); i++) {
            for (auto j = nums.size()-1; j > i; j--) {
                auto ii = i+1;
                auto jj = j-1;
                while (ii < jj) {
                    auto a = nums[i] * nums[j];
                    auto b = nums[ii] * nums[jj];
                    if (a == b) {
                        result++;
                        ii++;
                        jj--;
                    }
                    if (a > b) {
                        ii++;
                    }
                    if (a < b) {
                        jj--;
                    }
                }
            }
        }

        return result * 8;
    }
};
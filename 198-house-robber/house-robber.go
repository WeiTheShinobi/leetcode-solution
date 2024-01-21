func rob(nums []int) int {
    dp := make([]int, len(nums)+2)
    
    for i := 0; i < len(nums); i++ {
        dp[i+2] = max(dp[i+1], dp[i]+nums[i])
    }
    return dp[len(dp)-1]
}
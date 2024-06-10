func maximumLength(nums []int, k int) int {
	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, k+1)
	}

	result := 0
	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			for l := 0; l <= k; l++ {
				if nums[i] == nums[j] {
					dp[i][l] = max(dp[i][l], dp[j][l] + 1)
				} else if l > 0 {
					dp[i][l] = max(dp[i][l], dp[j][l-1] + 1)
				}
				result = max(result, dp[i][l])
			}
		}
	}

	return result + 1
}
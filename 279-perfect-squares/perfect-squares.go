func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
	}
	for i := 2; i < n; i++ {
		x := i * i
		if x > n {
			break
		}
		for j := x; j <= n; j++ {
			dp[j] = min(dp[j-x]+1, dp[j])
		}
	}

	return dp[n]
}
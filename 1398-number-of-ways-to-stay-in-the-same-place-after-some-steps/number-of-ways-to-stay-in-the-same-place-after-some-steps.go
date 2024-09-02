func numWays(steps int, arrLen int) int {
	if arrLen == 1 {
		return 1
	}
	mod := int(1e9 + 7)

	dp := make([][]int, min(arrLen, steps))
	for i := 0; i < min(arrLen, steps); i++ {
		dp[i] = make([]int, steps)
	}
	dp[0][steps-1] = 1
	dp[1][steps-1] = 1
	for j := steps - 2; j >= 0; j-- {
		for i := 0; i < min(arrLen, steps); i++ {
			dp[i][j] += dp[i][j+1]
			dp[i][j] %= mod
			if i > 0 {
				dp[i][j] += dp[i-1][j+1]
				dp[i][j] %= mod
			}
			if i < min(arrLen, steps)-1 {
				dp[i][j] += dp[i+1][j+1]
				dp[i][j] %= mod
			}
		}
	}
	return dp[0][0]
}
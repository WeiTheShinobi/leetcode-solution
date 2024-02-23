func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	inf := 10000*101 + 1
	dp := make([][]int, k+2)
	for i := 0; i < k+2; i++ {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	dp[0][src] = 0

	for t := 1; t < k+2; t++ {
		for _, flight := range flights {
			j, i, cost := flight[0], flight[1], flight[2]
			dp[t][i] = min(dp[t][i], dp[t-1][j]+cost)
		}
	}
	ans := inf
	for t := 1; t <= k+1; t++ {
		ans = min(ans, dp[t][dst])
	}
	if ans == inf {
		ans = -1
	}
	return ans
}
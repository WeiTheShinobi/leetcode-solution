func largestNumber(cost []int, target int) string {
	dp := make([]int, target+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MinInt
	}
	dp[0] = 0

	for _, n := range cost {
		for j := n; j <= target; j++ {
			dp[j] = max(dp[j], dp[j-n]+1)
		}
	}

	if dp[target] < 0 {
		return "0"
	}
	result := ""
	k := target
	for i := 9; i >= 1; i-- {
		for k >= cost[i-1] && dp[k-cost[i-1]]+1 == dp[k] {
			k -= cost[i-1]
			result += strconv.Itoa(i)
		}
	}
	return result
}
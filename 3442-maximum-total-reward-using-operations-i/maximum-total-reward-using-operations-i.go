func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	dp := make([][]bool, len(rewardValues)+1)
	for i := range dp {
		dp[i] = make([]bool, 4001)
	}

	dp[0][0] = true
	for i := 1; i < len(dp); i++ {
		v := rewardValues[i-1]
		for j := 0; j < 4000; j++ {
			if j >= v  && j - v < v {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	for i := 4000; i >= 0; i-- {
		if dp[len(dp)-1][i] {
			return i
		}
	}

	return 0
}
func maxTotalReward(rewardValues []int) int {
	m := slices.Max(rewardValues) * 2
	dp := make([]bool, m)
	dp[0] = true

	sort.Ints(rewardValues)
	for i := 0; i < len(rewardValues); i++ {
		for j := 0; j < rewardValues[i]; j++ {
			if dp[j] {
				dp[j+rewardValues[i]] = true
			}
		}
	}

	for i := len(dp) - 1; i >= 0; i-- {
		if dp[i] {
			return i
		}
	}
	return 0
}
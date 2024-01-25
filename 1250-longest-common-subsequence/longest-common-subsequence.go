func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)

	dp := make([]int, m+1)

	for i := 0; i < n; i++ {
		pre := 0
		for j := 0; j < m; j++ {
			if text1[i] == text2[j] {
				dp[j+1], pre = pre + 1, dp[j+1]
			} else {
        pre = dp[j+1]
				dp[j+1] = max(dp[j+1], dp[j])
			}
		}
	}

	return dp[m]

	//n, m := len(text1), len(text2)
	//
	//dp := make([][]int, n+1)
	//for i := 0; i < n+1; i++ {
	//	dp[i] = make([]int, m+1)
	//}
	//
	//for i := 0; i < n+1; i++ {
	//	for j := 0; j < m+1; j++ {
	//		if text1[i] == text2[j] {
	//			dp[i+1][j+1] = dp[i][j] + 1
	//		} else {
	//			dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
	//		}
	//	}
	//}
	//return dp[n][m]
}
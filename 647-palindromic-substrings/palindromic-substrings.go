func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	result := n
	for i := 1; i < n; i++ {
		dp[i-1][i] = s[i-1] == s[i]
		if dp[i-1][i] {
			result++
		}
	}
	for l := 2; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i+l
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			if dp[i][j] {
				result++
			}
		}
	}
	return result
}
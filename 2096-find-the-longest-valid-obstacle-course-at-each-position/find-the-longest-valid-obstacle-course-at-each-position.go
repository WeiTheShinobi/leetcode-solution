func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	result := make([]int, len(obstacles))
	var dp []int
	for i, obstacle := range obstacles {
		p := sort.SearchInts(dp, obstacle+1)
		if p >= len(dp) {
			dp = append(dp, obstacle)
		} else {
			dp[p] = obstacle
		}
		result[i] = p + 1
	}
	return result
}
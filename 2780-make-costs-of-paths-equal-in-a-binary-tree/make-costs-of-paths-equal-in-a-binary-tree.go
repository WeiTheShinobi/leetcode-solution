func minIncrements(n int, cost []int) int {
	result := 0
	for i := n / 2; i >= 1; i-- {
		left, right := cost[i*2-1], cost[i*2]
		if left > right {
			left, right = right, left
		}
		result += right - left
		cost[i-1] += right
	}
	return result
}
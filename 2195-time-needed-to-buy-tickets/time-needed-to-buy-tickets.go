func timeRequiredToBuy(tickets []int, k int) int {
	var result int
	n := tickets[k]

	for i := 0; i <= k; i++ {
		result += min(tickets[i], n)
	}
	for i := k + 1; i < len(tickets); i++ {
		result += min(tickets[i], n-1)
	}

	return result
}
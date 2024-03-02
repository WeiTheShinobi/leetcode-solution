func getMaximumConsecutive(coins []int) int {
	i := 0
	sort.Ints(coins)
	for _, coin := range coins {
		if coin > i+1 {
			break
		}
		i += coin
	}
	return i + 1
}
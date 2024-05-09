func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})

	var result int64
	for i := 0; i < k; i++ {
		result += int64(max(happiness[i]-i, 0))
	}

	return result
}
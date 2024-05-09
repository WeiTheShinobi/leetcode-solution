func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})

	var result int64
	for i := 0; i < k; i++ {
		v := int64(happiness[i] - i)
		if v < 0 {
			break
		}
		result += v
	}

	return result
}
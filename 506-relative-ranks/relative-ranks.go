func findRelativeRanks(score []int) []string {
	idx := make([]int, len(score))
	for i := 0; i < len(score); i++ {
		idx[i] = i
	}
	slices.SortFunc(idx, func(a, b int) int {
		return score[b] - score[a]
	})
	result := make([]string, len(score))
	for i := 0; i < len(score); i++ {
		if i == 0 {
			result[idx[i]] = "Gold Medal"
		} else if i == 1 {
			result[idx[i]] = "Silver Medal"
		} else if i == 2 {
			result[idx[i]] = "Bronze Medal"
		} else {
			result[idx[i]] = strconv.Itoa(i+1)
		}
	}

	return result
}

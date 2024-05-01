func wonderfulSubstrings(word string) int64 {
	var (
		result int64
		mask   int
	)

	count := [1024]int{1}
	for _, c := range word {
		mask ^= 1 << int(c-'a')
		result += int64(count[mask])

		for i := 1; i < 1024; i <<= 1 {
			result += int64(count[mask^i])
		}
		count[mask]++
	}

	return result
}
func wordCount(startWords []string, targetWords []string) int {
	has := make(map[int]bool)
	for _, word := range startWords {
		n := 0
		for i := 0; i < len(word); i++ {
			n |= 1 << (word[i] - 'a')
		}
		has[n] = true
	}

	var result int
	for _, word := range targetWords {
		target := 0
		for i := 0; i < len(word); i++ {
			target |= 1 << (word[i] - 'a')
		}

		for i := 0; i < len(word); i++ {
			if has[target^(1 << (word[i] - 'a'))] {
				result++
				break
			}
		}
	}
	return result
}
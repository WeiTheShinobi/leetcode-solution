func longestIdealString(s string, k int) int {
	if k >= 24 {
		return len(s)
	}
	var result int
	big := make([]int, 26)

	for _, c := range s {
		curr := 1

		for i := c - 'a' - int32(k); i <= c-'a'+int32(k) && i < 26; i++ {
			if i < 0 {
				i = 0
			}
			curr = max(curr, big[i]+1)
		}
		big[c-'a'] = curr
		result = max(result, curr)
	}

	return result
}
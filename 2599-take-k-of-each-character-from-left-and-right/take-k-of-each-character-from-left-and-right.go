func takeCharacters(s string, k int) int {
	count := make(map[uint8]int)

	for i := range s {
		count[s[i]]++
	}
	if count['a'] < k || count['b'] < k || count['c'] < k {
		return -1
	}

	result := len(s)
	for l, r := 0, 0; r < len(s); r++ {
		rc := s[r]
		count[rc]--
		for count[rc] < k {
			lc := s[l]
			count[lc]++
			l++
		}
		result = min(result, len(s)-(r-l+1))
	}
	return result
}
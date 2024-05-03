func kthPalindrome(queries []int, intLength int) []int64 {
	var result []int64
	base := 1
	for i := 0; i < (intLength-1)/2; i++ {
		base *= 10
	}

	for _, query := range queries {
		if query > 9*base {
			result = append(result, -1)
			continue
		}

		v := base + query - 1
		x := v
		if intLength&1 == 1 {
			x /= 10
		}
		for ; x > 0; x /= 10 {
			v = v*10 + x%10
		}
		result = append(result, int64(v))
	}
	return result
}
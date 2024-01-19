func numberOfWays(corridor string) int {
	const mod = 1e9 + 7
	pre, count, result := 0, 0, 1
	for i, c := range corridor {
		if c == 'S' {
			count++
			if count > 2 && count%2 == 1 {
				result *= i - pre
				result %= mod
			}
			pre = i
		}
	}

	if count < 2 || count%2 == 1 {
		return 0
	}
	return result
}
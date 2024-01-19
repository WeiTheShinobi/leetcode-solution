func numberOfWays(corridor string) int {
	const mod = 1e9 + 7
	result := 1
	chair := false
	chairNum := 0
	treeNum := 0
	totalChair := 0
	for _, c := range corridor {
		if chair {
			if c == 'P' {
				treeNum++
			} else {
				result *= treeNum + 1
        result %= mod
				treeNum = 0
				chair = false
			}
		}
		if !chair {
			if c == 'S' {
				chairNum++
        totalChair++
			}
			if chairNum == 2 {
				chair = true
				chairNum = 0
			}
		}
	}
	if totalChair%2 == 1 || totalChair < 2 {
		return 0
	}
	return result % mod
}
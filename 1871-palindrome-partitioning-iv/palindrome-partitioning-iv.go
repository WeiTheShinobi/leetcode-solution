func checkPartitioning(s string) bool {
	preCaclu := make([][]bool, len(s))
	for i := range preCaclu {
		preCaclu[i] = make([]bool, len(s))
		preCaclu[i][i] = true
	}

	for i := 1; i < len(s); i++ {
		preCaclu[i-1][i] = s[i] == s[i-1]
	}
	for l := 2; l < len(s); l++ {
		for i := l; i < len(s); i++ {
			preCaclu[i-l][i] = preCaclu[i-l+1][i-1] && s[i-l] == s[i]
		}
	}
	for i := 0; i < len(s); i++ {
		if preCaclu[0][i] {
			for j := i + 1; j < len(s)-1; j++ {
				if preCaclu[i+1][j] && preCaclu[j+1][len(s)-1] {
					return true
				}
			}
		}
	}

	return false
}
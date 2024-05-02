func platesBetweenCandles(s string, queries [][]int) []int {
	var result []int
	preSum := make([]int, len(s))
	left := make([]int, len(s))
	plates := 0
	l := -1
	for i := range s {
		if s[i] == '|' {
			for j := l + 1; j < i; j++ {
				left[j] = i
			}
			l = i
		} else {
			plates++
		}
		left[i] = l
		preSum[i] = plates
	}

	right := make([]int, len(s))
	plates = 0
	r := len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '|' {
			for j := r - 1; j > i; j-- {
				right[j] = i
			}
			r = i
		} else {
			plates++
		}
		right[i] = r
	}
	for _, query := range queries {
		l, r := left[query[0]], right[query[1]]
		if r < len(s) && l >= 0 {
			result = append(result, max(preSum[r]-preSum[l], 0))
		} else {
      result = append(result, 0)
    }
	}

	return result
}
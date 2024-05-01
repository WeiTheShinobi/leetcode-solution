func cycleLengthQueries(n int, queries [][]int) []int {
	var result []int
	for _, query := range queries {
		n1, n2, r := query[0], query[1], 0

		for n1 != n2 {
			if n1 > n2 {
				n1 /= 2
			} else if n2 > n1 {
				n2 /= 2
			}
			r++
		}

		result = append(result, r+1)
	}

	return result
}
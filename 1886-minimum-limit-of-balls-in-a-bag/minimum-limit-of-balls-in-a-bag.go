func minimumSize(nums []int, maxOperations int) int {
	maxN := 0
	for _, n := range nums {
		maxN = max(maxN, n)
	}

	return sort.Search(maxN, func(y int) bool {
		if y == 0 {
			return false
		}

		ops := 0
		for _, n := range nums {
			if n > y {
				ops += (n - 1) / y
			}
		}

		return ops <= maxOperations
	})
}
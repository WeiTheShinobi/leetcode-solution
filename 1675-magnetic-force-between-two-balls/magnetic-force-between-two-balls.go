func maxDistance(position []int, m int) int {
	check := func(step int) bool {
		pre := position[0]
		count := 1
		for _, p := range position[1:] {
			if p-pre >= step {
				count++
				pre = p
			}
		}
		return count >= m
	}

	sort.Ints(position)
	l, r, result := 0, position[len(position)-1]-position[0], 0

	for l <= r {
		step := (l + r) / 2
		if check(step) {
			l = step + 1
			result = step
		} else {
			r = step - 1
		}
	}

	return result
}
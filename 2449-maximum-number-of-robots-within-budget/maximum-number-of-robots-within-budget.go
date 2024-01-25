func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
	var q []int
	result, s, left := 0, int64(0), 0
	for right, ct := range chargeTimes {
		for len(q) > 0 && ct > chargeTimes[q[len(q)-1]] {
			q = q[:len(q)-1]
		}

		q = append(q, right)
		s += int64(runningCosts[right])
		for len(q) > 0 && int64(chargeTimes[q[0]])+int64(right-left+1)*s > budget {
			s -= int64(runningCosts[left])
			for len(q) > 0 && q[0] <= left {
				q = q[1:]
			}
			left++
		}

		result = max(result, right-left+1)
	}
	return result
}
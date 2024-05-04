func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	var result int

	q := initialBoxes

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			box := q[i]
			if status[box] == 0 {
				q = append(q, box)
				continue
			}

			result += candies[box]
			for _, k := range keys[box] {
				status[k] = 1
			}
			for _, cb := range containedBoxes[box] {
				q = append(q, cb)
			}
		}
		q = q[size:]
		stop := true
		for _, box := range q {
			if status[box] == 1 {
				stop = false
				break
			}
		}
		if stop {
			break
		}
	}

	return result
}
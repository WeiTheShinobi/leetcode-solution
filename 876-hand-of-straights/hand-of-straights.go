func isNStraightHand(hand []int, groupSize int) bool {
	type pair struct {
		n, count int
	}
	sort.Ints(hand)

	var count []pair
	curr := -1
	i := -1
	for _, n := range hand {
		if n != curr {
			count = append(count, pair{n: n, count: 1})
			curr = n
			i++
		} else {
			count[i].count++
		}
	}

	for i := 0; i < len(count); i++ {

		if count[i].count == 0 {
			continue
		}
    
		sub := count[i].count
		count[i].count -= sub
		for j := i+1; j < groupSize+i; j++ {
      if j >= len(count) {
        return false
      }
			if count[j-1].n+1 != count[j].n {
				return false
			}
			count[j].count -= sub
		}
	}

	for i := 0; i < len(count); i++ {
		if count[i].count != 0 {
			return false
		}
	}
	return true
}
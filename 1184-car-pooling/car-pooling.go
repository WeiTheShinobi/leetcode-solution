func carPooling(trips [][]int, capacity int) bool {
	d := make([]int, 1001)
	for _, t := range trips {
		d[t[1]] += t[0]
		d[t[2]] -= t[0]
	}
	
	s := 0
	for _, n := range d {
		s += n
		if s > capacity {
			return false
		}
	}
	return true
}
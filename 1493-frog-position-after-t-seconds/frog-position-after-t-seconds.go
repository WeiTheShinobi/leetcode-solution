func frogPosition(n int, edges [][]int, t int, target int) float64 {
	paths := make(map[int][]int)

	for _, edge := range edges {
		from, to := edge[0], edge[1]
		paths[from] = append(paths[from], to)
		paths[to] = append(paths[to], from)
	}

	type frog struct {
		curr  int
		last  int
		radio float64
	}

	frogs := []frog{{curr: 1, radio: 1, last: 1}}
	ratios := make([]float64, n+1)
	time := 0
	for time < t {
		size := len(frogs)
		for i := 0; i < size; i++ {
			f := frogs[i]
			nextPaths := paths[f.curr]
			before := len(frogs)
			for _, next := range nextPaths {
				if f.last == next {
					continue
				}
				radio := f.radio / float64(len(nextPaths)-1)
				if f.last == f.curr {
					radio = f.radio / float64(len(nextPaths))
				}
				frogs = append(frogs, frog{
					curr:  next,
					last:  f.curr,
					radio: radio,
				})
			}
			if len(frogs) == before {
				frogs = append(frogs, f)
			}
		}
		frogs = frogs[size:]
		time++
	}

	for _, f := range frogs {
		ratios[f.curr] += f.radio
	}
	return ratios[target]
}
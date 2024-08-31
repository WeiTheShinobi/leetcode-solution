func maximumTotalDamage(power []int) int64 {
	count := map[int]int{}
	for _, n := range power {
		count[n]++
	}
	a := make([]int, 0, len(count))
	for i := range count {
		a = append(a, i)
	}
	slices.Sort(a)

	f := make([]int, len(count)+1)
	j := 0
	for i, x := range a {
		for a[j] < x-2 {
			j++
		}
		f[i+1] = max(f[i], f[j]+x*count[x])
	}
	
	return int64(f[len(count)])
}
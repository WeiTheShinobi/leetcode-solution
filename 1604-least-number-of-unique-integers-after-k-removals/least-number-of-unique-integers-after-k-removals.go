func findLeastNumOfUniqueInts(arr []int, k int) int {
	counter := map[int]int{}

	for _, n := range arr {
		counter[n]++
	}

	var c []int
	for _, v := range counter {
		c = append(c, v)
	}

	sort.Ints(c)
	
	for i := 0; i < len(c); i++ {
		if k < c[i] {
			return len(c[i:])
		}
		k -= c[i]
	}
	
	return 0
}
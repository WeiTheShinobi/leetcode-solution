func minAnagramLength(s string) int {
	set := map[int]bool{}
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i == 0 {
			set[i] = true
			set[len(s)/i] = true
		}
	}
	var val []int
	for i := range set {
		val = append(val, i)
	}
	sort.Ints(val)
	if len(val) == 0 {
		return 1
	}

	isOk := func(v int) bool {
		counter := make([]int, 26)
		for i := 0; i < v; i++ {
			counter[s[i]-'a']++
		}

		for i := 0; i < len(s); i += v {
			counter2 := make([]int, 26)
			copy(counter2, counter)

			for j := i; j < i+v; j++ {
				counter2[s[j]-'a']--
			}
			for i := 0; i < 26; i++ {
				if counter2[i] != 0 {
					return false
				}
			}
		}
		return true
	}

	for i := 0; i < len(val); i++ {
		if isOk(val[i]) {
			return val[i]
		}
	}
	return val[len(val)-1]
}